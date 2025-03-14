package repositories

import (
	"github.com/dondany/go-projects/to-do-rest/models"
)

type TodoRepository interface {
	GetTodoLists() ([]*models.TodoList, error)
	GetTodoList(name string) (*models.TodoList, error)
	CreateTodoList(list models.TodoList) (*models.TodoList, error)
	UpdateTodoList(name string, list models.TodoList) (*models.TodoList, error)
	DeleteTodoList(name string) error

	CreateTodo(name string, todo models.Todo) (*models.Todo, error)
	UpdateTodo(name string, id int, todo models.Todo) (*models.Todo, error)
	DeleteTodo(name string, id int) error
}
