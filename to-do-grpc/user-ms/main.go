package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/dondany/go-projects/to-do-grpc/user-ms/db"
	"github.com/dondany/go-projects/to-do-grpc/user-ms/pb"
	"github.com/dondany/go-projects/to-do-grpc/user-ms/user"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("LISTEN_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connet to the db: %v", err)
	}
	defer db.Close()

	repo := user.NewPostgreUserRepository(db)
	server := user.NewServer(&repo)

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9000: %v", err)
	}
}