package models

import "time"

type Todo struct {
	ID        int    `json:"id"`
	ListID	int	`json:"listId"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	CreatedAt time.Time 
}

type TodoList struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Todos []Todo `json:"todos,omitempty"`
}
