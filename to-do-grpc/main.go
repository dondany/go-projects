package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/dondany/go-projects/to-do-grpc/pb"
	"github.com/dondany/go-projects/to-do-grpc/todo"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	db := connectDb()
	repo := todo.NewPostgreTodoRepository(db)
	server := todo.NewServer(repo)

	grpcServer := grpc.NewServer()

	pb.RegisterTodoServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9000: %v", err)
	}
}

func connectDb() *sql.DB {
	dsn := "host=localhost port=5432 user=admin password=admin dbname=db sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		panic(err)
	}
	return db
}

