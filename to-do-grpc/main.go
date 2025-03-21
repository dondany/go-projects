package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net"

	"github.com/dondany/go-projects/to-do-grpc/pb"
	"github.com/dondany/go-projects/to-do-grpc/todo"
	"github.com/dondany/go-projects/to-do-grpc/user"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	db := connectDb()
	defer db.Close()
	repo := todo.NewPostgreTodoRepository(db)
	server := todo.NewServer(repo)

	grpcServer := grpc.NewServer()

	userRepo := user.NewPostgreUserRepository(db)
	userServer := user.NewServer(&userRepo)

	pb.RegisterTodoServiceServer(grpcServer, server)
	pb.RegisterUserServiceServer(grpcServer, userServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9000: %v", err)
	}
}

func connectDb() *sql.DB {
	dsn := "host=localhost port=5432 user=admin password=admin dbname=db sslmode=disable"

	slog.Info("opening connection to db")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	slog.Info("succesfuly opened connection to db")

	slog.Info("pinging db")
	if err = db.Ping(); err != nil {
		db.Close()
		panic(err)
	}
	slog.Info("successfuly pinged db")
	return db
}

