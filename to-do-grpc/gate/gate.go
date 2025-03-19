package main

import (
	"log"
	"net/http"

	"github.com/dondany/go-projects/to-do-grpc/handlers"
	"github.com/dondany/go-projects/to-do-grpc/pb"
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

	service := pb.NewTodoServiceClient(conn)
	handler := handlers.NewTodoHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /lists", handler.GetTodoLists)
	mux.HandleFunc("GET /lists/{id}", handler.GetTodoList)
	mux.HandleFunc("POST /lists", handler.CreateTodoList)
	mux.HandleFunc("PATCH /lists/{id}", handler.UpdateTodoList)
	mux.HandleFunc("DELETE /lists/{id}", handler.DeleteTodoList)
	mux.HandleFunc("POST /lists/{id}/todos", handler.CreateTodo)
	mux.HandleFunc("PUT /lists/{list_id}/todos/{todo_id}", handler.UpdateTodo)
	mux.HandleFunc("DELETE /lists/{list_id}/todos/{todo_id}", handler.DeleteTodo)

	http.ListenAndServe(":8081", mux)
}