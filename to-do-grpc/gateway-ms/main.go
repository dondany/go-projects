package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/dondany/go-projects/to-do-grpc/gateway-ms/handlers"
	"github.com/dondany/go-projects/to-do-grpc/gateway-ms/metrics"
	"github.com/dondany/go-projects/to-do-grpc/gateway-ms/pb"
	"github.com/dondany/go-projects/to-do-grpc/gateway-ms/token"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)



func main() {
	userHost := os.Getenv("USER_MS_HOST")
	userPort := os.Getenv("USER_MS_PORT")
	var userConn *grpc.ClientConn
	userConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", userHost, userPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to user-ms: %s", err)
	}
	slog.Info("successfuly connected to user-ms")
	defer userConn.Close()

	todoHost := os.Getenv("TODO_MS_HOST")
	todoPort := os.Getenv("TODO_MS_PORT")
	var todoConn *grpc.ClientConn
	todoConn, err = grpc.NewClient(fmt.Sprintf("%s:%s", todoHost, todoPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to todoConn: %s", err)
	}
	slog.Info("successfuly connected to todo-ms")
	defer todoConn.Close()

	userService := pb.NewUserServiceClient(userConn)
	userHandler := handlers.NewUserHandler(userService)

	todoService := pb.NewTodoServiceClient(todoConn)
	todoHandler := handlers.NewTodoHandler(todoService)

	mux := http.NewServeMux()
	//prometheus metrics
	mux.Handle("/metrics", promhttp.Handler())

	//auth routes(unprotected)
	mux.HandleFunc("POST /user/register", userHandler.RegisterUser)
	mux.HandleFunc("POST /user/login", userHandler.LoginUser)

	//todolist routes
	mux.HandleFunc("GET /{userId}/lists", token.Protect(todoHandler.GetTodoLists))
	mux.HandleFunc("GET /{userId}/lists/{id}", token.Protect(todoHandler.GetTodoList))
	mux.HandleFunc("POST /{userId}/lists", token.Protect(todoHandler.CreateTodoList))
	mux.HandleFunc("PATCH /{userId}/lists/{id}", token.Protect(todoHandler.UpdateTodoList))
	mux.HandleFunc("DELETE /{userId}/lists/{id}", token.Protect(todoHandler.DeleteTodoList))

	//todolist/todo routes
	mux.HandleFunc("POST /{userId}/lists/{id}/todos", token.Protect(todoHandler.CreateTodo))
	mux.HandleFunc("PUT /{userId}/lists/{list_id}/todos/{todo_id}", token.Protect(todoHandler.UpdateTodo))
	mux.HandleFunc("DELETE /{userId}/lists/{list_id}/todos/{todo_id}", token.Protect(todoHandler.DeleteTodo))

	listenPort := os.Getenv("LISTEN_PORT")
	http.ListenAndServe(fmt.Sprintf(":%s", listenPort), metrics.GeneralMetricsMiddleware(mux))
}