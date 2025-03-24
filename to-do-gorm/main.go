package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Operation string

const (
	Add    = Operation("add")
	List   = Operation("list")
	Delete = Operation("delete")
	Toggle = Operation("toggle")
)

type Todo struct {
	gorm.Model
	Name      string
	Completed bool
}

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=admin password=admin dbname=gorm_db sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed db connection")
	}

	db.AutoMigrate(&Todo{})

	if len(os.Args) < 2 {
		fmt.Println("No parameters provided")
		return
	}

	operation := Operation(os.Args[1])
	switch operation {
	case List:
		listTodos(db)
	case Add:
		if len(os.Args) < 3 {
			fmt.Println("Todo name not provided")
			return
		}
		todo := Todo{
			Name: os.Args[2],
		}
		db.Create(&todo)
		listTodos(db)
	case Delete:
		if len(os.Args) < 3 {
			fmt.Println("Todo id not provided")
			return
		}
		db.Delete(&Todo{}, os.Args[2])
		listTodos(db)
	case Toggle:
		if len(os.Args) < 3 {
			fmt.Println("Todo id not provided")
			return
		}
		var todo Todo
		db.First(&todo, os.Args[2])
		db.Model(&todo).Update("Completed", !todo.Completed)
		listTodos(db)
	}
}

func listTodos(db *gorm.DB) {
	var todos []Todo
	db.Find(&todos)
	for _, t := range todos {
		completed := ""
		if t.Completed {
			completed = "X"
		}
		fmt.Printf("(%v) %v [%v]\n", t.ID, t.Name, completed)
	}
}
