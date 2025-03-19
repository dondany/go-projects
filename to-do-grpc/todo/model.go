package todo

import "time"

type Todo struct {
	ID        int32
	ListID int32
	Name      string
	Completed bool
	CreatedAt time.Time
}

type TodoList struct {
	ID int32
	Name string
	Todos []Todo
	CreatedAt time.Time
}