package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Operation string

const (
	Add    = Operation("add")
	List   = Operation("list")
	Delete = Operation("delete")
	Toggle = Operation("toggle")
)

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func (t Todo) String() string {
	var tick string
	if t.Completed {
		tick = "X"
	} else {
		tick = " "
	}
	return fmt.Sprintf("%v [%v]\t%v", t.ID, tick, t.Name)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No parameters provided")
		return
	}

	dataFile, err := os.OpenFile("data.json", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()
	var todos []Todo
	err = json.NewDecoder(dataFile).Decode(&todos)
	if err != nil {
		panic(err)
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
		todos = append(todos, CreateTodo(os.Args[2], len(todos)+1))
		ListTodos(todos)
		err = SaveTodos(todos, dataFile) 
		if err != nil {
			fmt.Println("Could not write to data file", err)
		}
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
		err = SaveTodos(todos, dataFile) 
		if err != nil {
			fmt.Println("Could not write to data file", err)
		}
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
		err = SaveTodos(todos, dataFile) 
		if err != nil {
			fmt.Println("Could not write to data file", err)
		}
	}
}

func ListTodos(todos []Todo) {
	for _, t := range todos {
		fmt.Println(t)
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

func SaveTodos(todos []Todo, dataFile *os.File) error {
	if _, err := dataFile.Seek(0, 0); err != nil {
		return err
	}
	if err := dataFile.Truncate(0); err != nil {
		fmt.Println("Error truncating file:", err)
		return err
	}
	return json.NewEncoder(dataFile).Encode(todos)
}