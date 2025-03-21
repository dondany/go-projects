package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/dondany/go-projects/to-do-grpc/pb"
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

	// c := pb.NewTodoServiceClient(conn)
	ctx := context.Background();

	userClient := pb.NewUserServiceClient(conn)

	// userRes, err := userClient.CreateUser(ctx, &pb.UserRequest{
	// 	Name: "Harry Potter",
	// 	Email: "hp@hgws.com",
	// 	Password: "pass",
	// })
	// if err != nil {
	// 	slog.Error("Error when calling Create User", "err", err)
	// 	os.Exit(-1)
	// }
	// slog.Info("Response from the server", "res", userRes)

	userRes, err := userClient.GetUserByEmail(ctx, &pb.UserEmail{Email: "hp@hgws.com"})
	if err != nil {
		slog.Error("Error when calling Get User", "err", err)
		os.Exit(-1)
	}
	slog.Info("Response from the server", "res", userRes)

	// loginRes, err := userClient.LoginUser(ctx, &pb.LoginRequest{Email: "hp@hgws.com", Password: "blabla"})
	// if err != nil {
	// 	slog.Error("Invalid credentials")
	// 	os.Exit(-1)
	// }
	// slog.Info("Response from the server", "res", loginRes)

	loginRes, err := userClient.LoginUser(ctx, &pb.LoginRequest{Email: "hp@hgws.com", Password: "pass"})
	if err != nil {
		slog.Error("Invalid credentials")
		os.Exit(-1)
	}
	slog.Info("Response from the server", "res", loginRes)

	// cRes, err := c.CreateTodoList(ctx, &pb.TodoListRequest{
	// 	Name: "My super new list",
	// })
	// if err != nil {
	// 	slog.Error("Error when calling Create Todo List", "err", err)
	// 	os.Exit(-1)
	// }
	// slog.Info("Response from the server", "res", cRes)

	// response, err := c.GetTodoLists(ctx, &pb.TodoListFilter{})
	// if err != nil {
	// 	slog.Error("Error when calling Get Todo List", "err", err)
	// 	os.Exit(-1)
	// }
	// slog.Info("Response from the server", "res", response)

	// uRes, err := c.UpdateTodoList(ctx, &pb.UpdateTodoListRequest{
	// 	Id: cRes.Id,
	// 	Name: "New Name",
	// })
	// if err != nil {
	// 	slog.Error("Error when calling Update Todo List", "err", err)
	// 	os.Exit(-1)
	// }
	// slog.Info("Response from the server", "res", uRes)
	// todoRes, err := c.CreateTodo(ctx, &pb.TodoRequest{
	// 	ListId: uRes.Id,
	// 	Name: "Do nothing",
	// })
	// if err != nil {
	// 	slog.Error("Error when calling Create Todo", "err", err)
	// 	os.Exit(-1)
	// }
	// slog.Info("Response from the server", "res", todoRes)

	// uTodoRes, err := c.UpdateTodo(ctx, &pb.TodoUpdateRequest{
	// 	Id: todoRes.Id,
	// 	Name: "Do nothing 2",
	// })
	// if err != nil {
	// 	slog.Error("Error when calling Update Todo", "err", err)
	// 	os.Exit(-1)
	// }
	// slog.Info("Response from the server", "res", uTodoRes)
}