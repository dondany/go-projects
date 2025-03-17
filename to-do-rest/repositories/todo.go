package repositories

import (
	"github.com/dondany/go-projects/to-do-rest/models"
)

type TodoRepository interface {
	GetTodoLists() ([]models.TodoList, error)
	GetTodoList(id int) (models.TodoList, error)
	CreateTodoList(list models.TodoList) (models.TodoList, error)
	UpdateTodoList(id int, list models.TodoList) (models.TodoList, error)
	DeleteTodoList(id int) error

	CreateTodo(listId int, todo models.Todo) (models.Todo, error)
	UpdateTodo(listId int, id int, todo models.Todo) (models.Todo, error)
	DeleteTodo(id int) error
}
