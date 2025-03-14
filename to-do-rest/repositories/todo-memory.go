package repositories

import (
	"errors"

	"github.com/dondany/go-projects/to-do-rest/models"
)

type TodoMemoryRepository struct {
	todoLists []*models.TodoList
}

func NewTodoMemoryRepository() TodoRepository {
	todos := []*models.Todo{
		{1, "Do smthg", false},
		{2, "Do smthg else", false},
	}
	todoLists := []*models.TodoList{
		{
			"My To Do List",
			todos,
		},
		{
			"Different To Do List",
			[]*models.Todo{},
		},
	}

	return &TodoMemoryRepository{todoLists: todoLists}
}

func (r *TodoMemoryRepository) GetTodoLists() ([]*models.TodoList, error) {
	return r.todoLists, nil
}

func (r *TodoMemoryRepository) GetTodoList(name string) (*models.TodoList, error) {
	for _, t := range r.todoLists {
		if t.Name == name {
			return t, nil
		}
	}
	return &models.TodoList{}, errors.New("could not find the requested list")
}

func (r *TodoMemoryRepository) CreateTodoList(list models.TodoList) (*models.TodoList, error) {
	list.Todos = []*models.Todo{}
	r.todoLists = append(r.todoLists, &list)
	return &list, nil
}

func (r *TodoMemoryRepository) UpdateTodoList(name string, list models.TodoList) (*models.TodoList, error) {
	l, err := r.GetTodoList(name)
	if err != nil {
		return nil, err
	}
	l.Name = list.Name
	return l, nil
}

func (r *TodoMemoryRepository) DeleteTodoList(name string) error {
	for i, l := range r.todoLists {
		if l.Name == name {
			r.todoLists = append(r.todoLists[:i], r.todoLists[i+1:]...)
			return nil
		}
	}
	return errors.New("could not find the requested todo list to be deleted")
}

func (r *TodoMemoryRepository) CreateTodo(name string, todo models.Todo) (*models.Todo, error) {
	l, err := r.GetTodoList(name)
	if err != nil {
		return nil, err
	}
	l.Todos = append(l.Todos, &todo)
	return &todo, nil
}

func (r *TodoMemoryRepository) UpdateTodo(name string, id int, todo models.Todo) (*models.Todo, error) {
	l, err := r.GetTodoList(name)
	if err != nil {
		return nil, err
	}
	for _, t := range l.Todos {
		if t.ID == id {
			t.Completed = todo.Completed
			t.Name = todo.Name
			return t, nil
		}
	}
	return nil, errors.New("could not find the requested todo to be updated")
}

func (r *TodoMemoryRepository) DeleteTodo(name string, id int) error {
	l, err := r.GetTodoList(name)
	if err != nil {
		return err
	}
	for i, t := range l.Todos {
		if t.ID == id {
			l.Todos = append(l.Todos[:i], l.Todos[i+1:]...)
			return nil
		}
	}
	return errors.New("could not find the requested todo to be deleted")
}
