package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dondany/go-projects/to-do-grpc/gateway-ms/model"
	"github.com/dondany/go-projects/to-do-grpc/gateway-ms/pb"
	"github.com/dondany/go-projects/to-do-grpc/gateway-ms/token"
)

type UserHandler struct {
	service pb.UserServiceClient
}

func NewUserHandler(service pb.UserServiceClient) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (h UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	slog.Info("POST /users/register")

	var inputUser model.User
	err := json.NewDecoder(r.Body).Decode(&inputUser)
	if err != nil {
		fmt.Fprintf(w, `{"error": "Wrong request}`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userReq := &pb.UserRequest{
		Email: inputUser.Email,
		Name: inputUser.Name,
		Password: inputUser.Password,
	}

	response, err := h.service.CreateUser(context.Background(), userReq)
	if err != nil {
		http.Error(w, `{"error": "invalid credentials}`, http.StatusUnauthorized)
	}

	tokenString, err := token.Create(response.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, `{"error": "failed token generations}`, http.StatusInternalServerError)
	}

	registeredUser := model.User{
		ID: response.Id,
		Email: response.Email,
		Name: response.Name,
		CreatedAt: response.CreatedAt.AsTime(),
		Token: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(registeredUser)
}

func (h UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	slog.Info("POST /users/login")

	var inputUser model.UserLogin
	err := json.NewDecoder(r.Body).Decode(&inputUser)
	if err != nil {
		fmt.Fprintf(w, `{"error": "Wrong request}`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userReq := &pb.LoginRequest{
		Email: inputUser.Email,
		Password: inputUser.Password,
	}

	response, err := h.service.LoginUser(context.Background(), userReq)
	if err != nil {
		http.Error(w, `{"error": "invalid credentials}`, http.StatusUnauthorized)
	}

	tokenString, err := token.Create(response.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, `{"error": "failed token generations}`, http.StatusInternalServerError)
	}

	user := model.User{
		ID: response.Id,
		Email: response.Email,
		Name: response.Name,
		CreatedAt: response.CreatedAt.AsTime(),
		Token: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}