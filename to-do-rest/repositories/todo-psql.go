package repositories

import (
	"database/sql"
	"fmt"

	"github.com/dondany/go-projects/to-do-rest/models"
)

type TodoPostgreSqlRepository struct {
	db *sql.DB
}

func NewTodoPostgreSqlRepository(db *sql.DB) TodoRepository {
	return TodoPostgreSqlRepository{
		db: db,
	}
}

func (r TodoPostgreSqlRepository) GetTodoLists() ([]models.TodoList, error) {
	rows, err := r.db.Query("select id, name from lists")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []models.TodoList
	for rows.Next() {
		var list models.TodoList
		if err := rows.Scan(&list.ID, &list.Name); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}

func (r TodoPostgreSqlRepository) GetTodoList(id int) (models.TodoList, error) {
	query := `
	select tl.id, tl.name, t.id, t.name, t.completed
	from lists tl
	left join todos t ON tl.id = t.list_id
	where tl.id = $1
	`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return models.TodoList{}, err
	}

	var todoList models.TodoList
	todoList.Todos = []models.Todo{}
	var todoID sql.NullInt64
	var todoName sql.NullString
	var todoCompleted sql.NullBool

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todoList.ID, &todoList.Name, &todoID, &todoName, &todoCompleted); err != nil {
			fmt.Println(err)
			return models.TodoList{}, err
		}
		if todoID.Valid && todoName.Valid && todoCompleted.Valid {
			todo.ID = int(todoID.Int64)
			todo.Name = todoName.String
			todo.Completed = todoCompleted.Bool
			todoList.Todos = append(todoList.Todos, todo)
		}
	}

	return todoList, nil
}

func (r TodoPostgreSqlRepository) CreateTodoList(list models.TodoList) (models.TodoList, error) {
	var newList models.TodoList
	err := r.db.QueryRow("insert into lists (name) values ($1) returning id, name", list.Name).Scan(&newList.ID, &newList.Name)
	if err != nil {
		fmt.Println(err)
		return models.TodoList{}, err
	}
	return newList, nil
}

func (r TodoPostgreSqlRepository) UpdateTodoList(id int, list models.TodoList) (models.TodoList, error) {
	var updatedList models.TodoList

	// Update the TodoList and return the updated record
	err := r.db.QueryRow(
		"update lists set name = $1 WHERE id = $2 RETURNING id, name",
		list.Name, id,
	).Scan(&updatedList.ID, &updatedList.Name)
	if err != nil {
		fmt.Println(err)
		return models.TodoList{}, err
	}
	return updatedList, nil
}

func (r TodoPostgreSqlRepository) DeleteTodoList(id int) error {
	var targetId int
	err := r.db.QueryRow("delete from lists where id = $1 returning id", id).Scan(&targetId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("list with id %v not found", id)
		}
	}
	return nil
}

func (r TodoPostgreSqlRepository) CreateTodo(listId int, todo models.Todo) (models.Todo, error) {
	var newTodo models.Todo
	err := r.db.QueryRow("insert into todos (name, completed, list_id) values ($1, $2, $3) returning id, name, completed", todo.Name, todo.Completed, listId).Scan(&newTodo.ID, &newTodo.Name, &newTodo.Completed)
	if err != nil {
		fmt.Println(err)
		return models.Todo{}, err
	}
	return newTodo, nil
}

func (r TodoPostgreSqlRepository) UpdateTodo(listId int, id int, todo models.Todo) (models.Todo, error) {
	var updatedTodo models.Todo
	err := r.db.QueryRow(
		"update todos set name = $1, completed = $2 WHERE id = $3 RETURNING id, name, completed",
		todo.Name, todo.Completed, id,
	).Scan(&updatedTodo.ID, &updatedTodo.Name, &updatedTodo.Completed)
	if err != nil {
		fmt.Println(err)
		return models.Todo{}, err
	}
	return updatedTodo, nil
}

func (r TodoPostgreSqlRepository) DeleteTodo(id int) error {
	var targetId int
	err := r.db.QueryRow("delete from todos where id = $1 returning id", id).Scan(&targetId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("list with id %v not found", id)
		}
	}
	return nil
}
