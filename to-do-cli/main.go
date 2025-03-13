package main

import (
	"fmt"
	"os"
	"strconv"
)

type Operation string

const (
	Add = Operation("add")
	List = Operation("list")
	Delete = Operation("delete")
	Toggle = Operation("toggle")
)

type Todo struct {
	ID int
	Name string
	Completed bool
}

func (t Todo)String() string {
	return fmt.Sprintf("%v\t%v\t%v", t.ID, t.Name, t.Completed)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No parameters provided")
		return
	}

	todos := []Todo{
		{
			1,
			"Do something",
			false,
		},
		{
			24,
			"Do nothing",
			true,
		},
	}

	operation := Operation(os.Args[1])

	switch operation {
	case List:
		ListTodos(todos)
	case Add:
		if len(os.Args) < 3 {
			fmt.Println("Todo name not provided")
			return
		}
		todos = append(todos, CreateTodo(os.Args[2], len(todos)))
		ListTodos(todos)
	case Delete:
		if len(os.Args) < 3 {
			fmt.Println("Todo id not provided")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Wrong id provided. Has to be a number")
			return
		}
		todos, ok := DeleteTodo(todos, id)
		if !ok {
			fmt.Println("No todo found with the given id")
			return
		}
		ListTodos(todos)
	case Toggle:
		if len(os.Args) < 3 {
			fmt.Println("Todo id not provided")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Wrong id provided. Has to be a number")
			return
		}
		ok := ToggleTodo(todos, id)
		if !ok {
			fmt.Println("No todo found with the given id")
			return
		}
		ListTodos(todos)
	}

}

func ListTodos(todos []Todo) {
	for _, t := range todos {
		fmt.Println(t);
	}
}

func CreateTodo(name string, id int) Todo {
	return Todo{
		id,
		name,
		false,
	}
}

func DeleteTodo(todos []Todo, id int) ([]Todo, bool) {
	result := make([]Todo, 0, len(todos))
	var found bool
	for _, t := range todos {
		if t.ID != id {
			result = append(result, t)
		} else {
			found = true
			break
		}
	}
	return result, found
}

func ToggleTodo(todos []Todo, id int) bool {
	var found bool
	for i, t := range todos {
		if t.ID == id {
			todos[i].Completed = !todos[i].Completed
			found = true
			break
		}
	}
	return found
}