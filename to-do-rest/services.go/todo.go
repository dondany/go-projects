package services

import (
	"github.com/dondany/go-projects/to-do-rest/models"
	"github.com/dondany/go-projects/to-do-rest/repositories"
)

type TodoService interface {
	GetTodoLists() ([]*models.TodoList, error)
	GetTodoList(name string) (*models.TodoList, error)
	CreateTodoList(list models.TodoList) (*models.TodoList, error)
	UpdateTodoList(name string, list models.TodoList) (*models.TodoList, error)
	DeleteTodoList(name string) error

	CreateTodo(name string, todo models.Todo) (*models.Todo, error)
	UpdateTodo(name string, id int, todo models.Todo) (*models.Todo, error)
	DeleteTodo(name string, id int) error
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo}
}

func (t *todoService) GetTodoLists() ([]*models.TodoList, error) {
	return t.repo.GetTodoLists()
}

func (t *todoService) GetTodoList(name string) (*models.TodoList, error) {
	return t.repo.GetTodoList(name)
}

func (t *todoService) CreateTodoList(list models.TodoList) (*models.TodoList, error) {
	return t.repo.CreateTodoList(list)
}

func (t *todoService) DeleteTodoList(name string) error {
	return t.repo.DeleteTodoList(name)
}

func (t *todoService) UpdateTodoList(name string, list models.TodoList) (*models.TodoList, error) {
	return t.repo.UpdateTodoList(name, list)
}

func (t *todoService) CreateTodo(name string, todo models.Todo) (*models.Todo, error) {
	return t.repo.CreateTodo(name, todo)
}

func (t *todoService) UpdateTodo(name string, id int, todo models.Todo) (*models.Todo, error) {
	return t.repo.UpdateTodo(name, id, todo)
}

func (t *todoService) DeleteTodo(name string, id int) error {
	return t.repo.DeleteTodo(name, id)
}
