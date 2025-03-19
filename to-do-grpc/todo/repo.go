package todo

import (
	"database/sql"
	"fmt"
	"time"
)

type TodoRepository interface {
	CreateTodoList(TodoList) (TodoList, error)
	GetTodoList(listId int32) (TodoList, error)
	GetTodoLists() ([]TodoList, error)
	UpdateTodoList(TodoList) (TodoList, error)
	DeleteTodoList(listId int32) error

	CreateTodo(todo Todo) (Todo, error)
	UpdateTodo(todo Todo) (Todo, error)
	DeleteTodo(id int32) error
}

type PostgreTodoRepository struct {
	db *sql.DB
}

func NewPostgreTodoRepository(db *sql.DB) PostgreTodoRepository {
	return PostgreTodoRepository{
		db: db,
	}
}

func (r PostgreTodoRepository) CreateTodoList(list TodoList) (TodoList, error) {
	now := time.Now()
	var newList TodoList
	err := r.db.QueryRow("insert into lists (name, created_at) values ($1, $2) returning id, name, created_at", list.Name, now).Scan(&newList.ID, &newList.Name, &newList.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return TodoList{}, err
	}
	return newList, nil
}

func (r PostgreTodoRepository) GetTodoLists() ([]TodoList, error) {
	rows, err := r.db.Query("select id, name from lists")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []TodoList
	for rows.Next() {
		var list TodoList
		if err := rows.Scan(&list.ID, &list.Name); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}

func (r PostgreTodoRepository) GetTodoList(id int32) (TodoList, error) {
	query := `
	select tl.id, tl.name, t.id, t.name, t.completed
	from lists tl
	left join todos t ON tl.id = t.list_id
	where tl.id = $1
	`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return TodoList{}, err
	}

	var todoList TodoList
	todoList.Todos = []Todo{}
	var todoID sql.NullInt32
	var todoName sql.NullString
	var todoCompleted sql.NullBool

	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todoList.ID, &todoList.Name, &todoID, &todoName, &todoCompleted); err != nil {
			fmt.Println(err)
			return TodoList{}, err
		}
		if todoID.Valid && todoName.Valid && todoCompleted.Valid {
			todo.ID = todoID.Int32
			todo.Name = todoName.String
			todo.Completed = todoCompleted.Bool
			todoList.Todos = append(todoList.Todos, todo)
		}
	}

	return todoList, nil
}

func (r PostgreTodoRepository) UpdateTodoList(list TodoList) (TodoList, error) {
	var updatedList TodoList
	err := r.db.QueryRow(
		"update lists set name = $1 WHERE id = $2 RETURNING id, name",
		list.Name, list.ID,
	).Scan(&updatedList.ID, &updatedList.Name)
	if err != nil {
		fmt.Println(err)
		return TodoList{}, err
	}
	return updatedList, nil
}

func (r PostgreTodoRepository) DeleteTodoList(id int32) error {
	var targetId int32
	err := r.db.QueryRow("delete from lists where id = $1 returning id", id).Scan(&targetId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("list with id %v not found", id)
		}
	}
	return nil
}

func (r PostgreTodoRepository) CreateTodo(todo Todo) (Todo, error) {
	now := time.Now()
	var newTodo Todo
	err := r.db.QueryRow("insert into todos (name, completed, list_id, created_at) values ($1, $2, $3, $4) returning id, name, completed, created_at", todo.Name, todo.Completed, todo.ListID, now).Scan(&newTodo.ID, &newTodo.Name, &newTodo.Completed, &newTodo.CreatedAt)
	if err != nil {
		return Todo{}, err
	}
	return newTodo, nil
}

func (r PostgreTodoRepository) UpdateTodo(todo Todo) (Todo, error) {
	var updatedTodo Todo
	err := r.db.QueryRow(
		"update todos set name = $1, completed = $2 WHERE id = $3 RETURNING id, name, completed",
		todo.Name, todo.Completed, todo.ID,
	).Scan(&updatedTodo.ID, &updatedTodo.Name, &updatedTodo.Completed)
	if err != nil {
		return Todo{}, err
	}
	return updatedTodo, nil
}

func (r PostgreTodoRepository) DeleteTodo(id int32) error {
	var targetId int32
	err := r.db.QueryRow("delete from todos where id = $1 returning id", id).Scan(&targetId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("list with id %v not found", id)
		}
	}
	return nil
}