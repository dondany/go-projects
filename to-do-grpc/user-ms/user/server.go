package user

import (
	"context"

	"github.com/dondany/go-projects/to-do-grpc/user-ms/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	repo UserRepository
	pb.UnimplementedUserServiceServer
}

func NewServer(repo UserRepository) *Server {
	return &Server{repo: repo}
}

func (s *Server) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}
	user := User{
		Name: req.Name,
		Email: req.Email,
		Password: string(hashedPass),
	}
	newUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create user: %v", err)
	}

	response := &pb.UserResponse{
		Id: newUser.ID,
		Name: newUser.Name,
		Email: newUser.Email,
		CreatedAt: timestamppb.New(newUser.CreatedAt),
	}
	return response, nil
}

func (s *Server) GetUserByEmail(ctx context.Context, req *pb.UserEmail) (*pb.UserResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to fetch user: %v", err)
	}

	response := &pb.UserResponse{
		Id: user.ID,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
	return response, nil
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.UserResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to fetch user: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "invalid credentatials")
	}

	response := &pb.UserResponse{
		Id: user.ID,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
	return response, nil
}