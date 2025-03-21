package main

import (
	"log"
	"net/http"

	"github.com/dondany/go-projects/to-do-grpc/handlers"
	"github.com/dondany/go-projects/to-do-grpc/pb"
	"github.com/dondany/go-projects/to-do-grpc/token"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	todoService := pb.NewTodoServiceClient(conn)
	todoHandler := handlers.NewTodoHandler(todoService)

	userService := pb.NewUserServiceClient(conn)
	userHandler := handlers.NewUserHandler(userService)

	mux := http.NewServeMux()

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

	http.ListenAndServe(":8081", mux)
}