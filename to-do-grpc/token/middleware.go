package token

import (
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // Bearer keyword is missing
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		claims, err := Verify(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		uriUserId := r.PathValue("userId")
		if uriUserId != claims["userId"] {
			http.Error(w, "forbidden! no access", http.StatusForbidden)
			return
		}

		md := metadata.Pairs("user_id", uriUserId)
		ctx := metadata.NewOutgoingContext(r.Context(), md)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func Protect(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return JWTMiddleware(handlerFunc).ServeHTTP
}