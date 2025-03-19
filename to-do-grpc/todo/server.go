package todo

import (
	"context"

	"github.com/dondany/go-projects/to-do-grpc/pb"
	"google.golang.org/grpc/codes"
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
	
	list, err := s.repo.CreateTodoList(TodoList{Name: req.Name})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create a todolist: %v", err)
	}

	response := pb.TodoListResponse{
		Id: list.ID,
		Name: list.Name,
	}

	return &response, nil
}

func (s *Server) GetTodoList(ctx context.Context, id *pb.ID) (*pb.TodoListResponse, error) {
	list, err := s.repo.GetTodoList(id.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get a todolist: %v", err)
	}
	
	response := pb.TodoListResponse{
		Id: list.ID,
		Name: list.Name,
	}

	return &response, nil
}

func (s *Server) GetTodoLists(ctx context.Context, filter *pb.TodoListFilter) (*pb.TodoListsResponse, error) {
	lists, err := s.repo.GetTodoLists()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get a todolist: %v", err)
	}

	responseLists := make([]*pb.TodoListResponse, len(lists), cap(lists))
	for _, l := range lists {
		list := pb.TodoListResponse{
			Id: l.ID,
			Name: l.Name,
			CreatedAt: timestamppb.New(l.CreatedAt),
		}
		responseLists = append(responseLists, &list)
	}

	response := pb.TodoListsResponse{
		Lists: responseLists,
	}
	return &response, nil
}

func (s *Server) UpdateTodoList(ctx context.Context, req *pb.UpdateTodoListRequest) (*pb.TodoListResponse, error) {
	list := TodoList{
		ID: req.Id,
		Name: req.Name,
	}
	updatedList, err := s.repo.UpdateTodoList(list)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update a todolist: %v", err)
	}

	response := pb.TodoListResponse{
		Id: updatedList.ID,
		Name: updatedList.Name,
		CreatedAt: timestamppb.New(updatedList.CreatedAt),
	}

	return &response, nil
}

func (s *Server) DeleteTodoList(ctx context.Context, id *pb.ID) (*emptypb.Empty, error) {
	err := s.repo.DeleteTodoList(id.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete a todolist: %v", err)
	}
	return nil, nil
}

func (s *Server) CreateTodo(ctx context.Context, todo *pb.TodoRequest) (*pb.TodoResponse, error) {
	newTodo, err := s.repo.CreateTodo(Todo{
		ID: todo.Id,
		ListID: todo.ListId,
		Name: todo.Name,
		Completed: false,
	})
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
	updatedTodo, err := s.repo.UpdateTodo(Todo{
		ID: todo.Id,
		Name: todo.Name,
		Completed: false,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update todo: %v", err)
	}
	
	response := pb.TodoResponse{
		Id: updatedTodo.ID,
		ListId: updatedTodo.ListID,
		Name: updatedTodo.Name,
		Completed: false,
	}
	return &response, nil
}

func (s *Server) DeleteTodo(ctx context.Context, id *pb.ID) (*emptypb.Empty, error) {
	err := s.repo.DeleteTodo(id.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete todo: %v", err)
	}
	return nil, nil
}