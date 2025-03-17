package services

import (
	"github.com/dondany/go-projects/to-do-rest/models"
	"github.com/dondany/go-projects/to-do-rest/repositories"
)

type TodoService interface {
	GetTodoLists() ([]models.TodoList, error)
	GetTodoList(id int) (models.TodoList, error)
	CreateTodoList(list models.TodoList) (models.TodoList, error)
	UpdateTodoList(id int, list models.TodoList) (models.TodoList, error)
	DeleteTodoList(id int) error

	CreateTodo(listId int, todo models.Todo) (models.Todo, error)
	UpdateTodo(listId int, id int, todo models.Todo) (models.Todo, error)
	DeleteTodo(id int) error
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo}
}

func (t *todoService) GetTodoLists() ([]models.TodoList, error) {
	return t.repo.GetTodoLists()
}

func (t *todoService) GetTodoList(id int) (models.TodoList, error) {
	return t.repo.GetTodoList(id)
}

func (t *todoService) CreateTodoList(list models.TodoList) (models.TodoList, error) {
	return t.repo.CreateTodoList(list)
}

func (t *todoService) DeleteTodoList(id int) error {
	return t.repo.DeleteTodoList(id)
}

func (t *todoService) UpdateTodoList(id int, list models.TodoList) (models.TodoList, error) {
	return t.repo.UpdateTodoList(id, list)
}

func (t *todoService) CreateTodo(listId int, todo models.Todo) (models.Todo, error) {
	return t.repo.CreateTodo(listId, todo)
}

func (t *todoService) UpdateTodo(listId int, id int, todo models.Todo) (models.Todo, error) {
	return t.repo.UpdateTodo(listId, id, todo)
}

func (t *todoService) DeleteTodo(id int) error {
	return t.repo.DeleteTodo(id)
}
