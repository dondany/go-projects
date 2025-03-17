package models

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type TodoList struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Todos []*Todo `json:"todos,omitempty"`
}
