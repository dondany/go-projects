package todo

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/dondany/go-projects/to-do-grpc/to-do-ms/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	repo TodoRepository
	pb.UnimplementedTodoServiceServer
}

func NewServer(repo TodoRepository) *Server {
	return &Server{repo: repo}
}

func (s *Server) CreateTodoList(ctx context.Context, req *pb.TodoListRequest) (*pb.TodoListResponse, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name cannot be empty")
	}
	
	list, err := s.repo.CreateTodoList(TodoList{Name: req.Name, UserID: req.UserId})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create a todolist: %v", err)
	}
	response := pb.TodoListResponse{
		Id: list.ID,
		Name: list.Name,
		UserId: list.UserID,
		CreatedAt: timestamppb.New(list.CreatedAt),
	}

	return &response, nil
}

func (s *Server) GetTodoList(ctx context.Context, id *pb.ID) (*pb.TodoListResponse, error) {
	list, err := s.repo.GetTodoList(id.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get a todolist: %v", err)
	}

	err = verifyAuthorization(ctx, list.UserID)
	if err != nil {
		return nil, err
	}

	todos := make([]*pb.TodoResponse, len(list.Todos))
	for i, t := range list.Todos {
		todos[i] = &pb.TodoResponse{
			Id: t.ID,
			ListId: t.ListID,
			Name: t.Name,
			Completed: t.Completed,
			CreatedAt: timestamppb.New(t.CreatedAt),
		}
	}
	
	response := pb.TodoListResponse{
		Id: list.ID,
		Name: list.Name,
		UserId: list.UserID,
		Todos: todos,
		CreatedAt: timestamppb.New(list.CreatedAt),
	}
	return &response, nil
}

func (s *Server) GetTodoLists(ctx context.Context, filter *pb.TodoListFilter) (*pb.TodoListsResponse, error) {
	lists, err := s.repo.GetTodoLists(filter.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get a todolist: %v", err)
	}

	responseLists := make([]*pb.TodoListResponse, len(lists))

	for i, l := range lists {
		list := pb.TodoListResponse{
			Id: l.ID,
			Name: l.Name,
			UserId: l.UserID,
			CreatedAt: timestamppb.New(l.CreatedAt),
		}
		responseLists[i] = &list
	}

	response := pb.TodoListsResponse{
		Lists: responseLists,
	}
	return &response, nil
}

func (s *Server) UpdateTodoList(ctx context.Context, req *pb.UpdateTodoListRequest) (*pb.TodoListResponse, error) {
	list, err := s.repo.GetTodoList(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get a todolist: %v", err)
	}

	err = verifyAuthorization(ctx, list.UserID)
	if err != nil {
		return nil, err
	}

	list.Name = req.Name

	updatedList, err := s.repo.UpdateTodoList(list)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update a todolist: %v", err)
	}

	response := pb.TodoListResponse{
		Id: updatedList.ID,
		Name: updatedList.Name,
		UserId: updatedList.UserID,
		CreatedAt: timestamppb.New(updatedList.CreatedAt),
	}

	return &response, nil
}

func (s *Server) DeleteTodoList(ctx context.Context, id *pb.ID) (*emptypb.Empty, error) {
	list, err := s.repo.GetTodoList(id.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get a todolist: %v", err)
	}

	err = verifyAuthorization(ctx, list.UserID)
	if err != nil {
		return nil, err
	}

	err = s.repo.DeleteTodoList(id.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete a todolist: %v", err)
	}
	return nil, nil
}

func (s *Server) CreateTodo(ctx context.Context, todo *pb.TodoRequest) (*pb.TodoResponse, error) {
	list, err := s.repo.GetTodoList(todo.ListId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get a todolist: %v", err)
	}

	err = verifyAuthorization(ctx, list.UserID)
	if err != nil {
		return nil, err
	}

	newTodo, err := s.repo.CreateTodo(Todo{
		ID: todo.Id,
		ListID: todo.ListId,
		Name: todo.Name,
		Completed: false,
	})
	slog.Info("Created", "Todo", newTodo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create todo: %v", err)
	}
	
	response := pb.TodoResponse{
		Id: newTodo.ID,
		ListId: newTodo.ListID,
		Name: newTodo.Name,
		Completed: false,
	}
	return &response, nil
}

func (s *Server) UpdateTodo(ctx context.Context, todo *pb.TodoUpdateRequest) (*pb.TodoResponse, error) {
	userId, err := s.repo.GetTodoUserId(todo.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "todo not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update todo: %v", err)
	}

	err = verifyAuthorization(ctx, userId);
	if err != nil {
		return nil, err
	}

	updatedTodo, err := s.repo.UpdateTodo(Todo{
		ID: todo.Id,
		Name: todo.Name,
		Completed: todo.Completed,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update todo: %v", err)
	}
	
	response := pb.TodoResponse{
		Id: updatedTodo.ID,
		ListId: updatedTodo.ListID,
		Name: updatedTodo.Name,
		Completed: todo.Completed,
	}
	return &response, nil
}

func (s *Server) DeleteTodo(ctx context.Context, id *pb.ID) (*emptypb.Empty, error) {
	userId, err := s.repo.GetTodoUserId(id.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "todo not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update todo: %v", err)
	}

	err = verifyAuthorization(ctx, userId);
	if err != nil {
		return nil, err
	}

	err = s.repo.DeleteTodo(id.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete todo: %v", err)
	}
	return nil, nil
}

//should probably be moved to interceptor
func verifyAuthorization(ctx context.Context, userId int32) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "failed to get a todolist, missing metadata")
	}
	userIDs := md.Get("user_id")
	if len(userIDs) == 0 {
		return status.Errorf(codes.InvalidArgument, "failed to get a todolist, missing user_id in metadata")
	}
	userID := userIDs[0]
	if fmt.Sprint(userId) != userID {
		return status.Errorf(codes.PermissionDenied, "unauthorized access")
	}
	return nil
}