package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/dondany/go-projects/to-do-rest/models"
	"github.com/dondany/go-projects/to-do-rest/repositories"
	"github.com/dondany/go-projects/to-do-rest/services"

	_ "github.com/lib/pq"
)

func main() {
	mux := http.NewServeMux()

	dsn := "host=localhost port=5432 user=admin password=admin dbname=db sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		panic(err)
	}

	psqlRepo := repositories.NewTodoPostgreSqlRepository(db)
	service := services.NewTodoService(psqlRepo)

	mux.HandleFunc("GET /lists", getTodoLists(service))
	mux.HandleFunc("GET /lists/{id}", getTodoList(service))
	mux.HandleFunc("POST /lists", createTodoList(service))
	mux.HandleFunc("PATCH /lists/{id}", updateTodoList(service))
	mux.HandleFunc("DELETE /lists/{id}", deleteTodoList(service))

	mux.HandleFunc("POST /lists/{id}/todos", createTodo(service))
	mux.HandleFunc("PUT /lists/{list_id}/todos/{todo_id}", updateTodo(service))
	mux.HandleFunc("DELETE /lists/{list_id}/todos/{todo_giid}", deleteTodo(service))

	err = http.ListenAndServe(":8081", mux)
	fmt.Println(err)
}

func getTodoLists(service services.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("GET /lists")
		lists, err := service.GetTodoLists()
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not fetch lists"))
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(lists)
	}
}

func getTodoList(service services.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		slog.Info("GET /lists/" + idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong id format"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		list, err := service.GetTodoList(id)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not fetch a list"))
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(list)
	}
}

func createTodoList(service services.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("POST /lists/")

		var list models.TodoList
		err := json.NewDecoder(r.Body).Decode(&list)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong request"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		createdList, err := service.CreateTodoList(list)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not create a Todo List"))
			return
		}

		json.NewEncoder(w).Encode(createdList)
	}
}

func updateTodoList(service services.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		slog.Info("PATCH /lists/" + idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong id format"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var list models.TodoList
		err = json.NewDecoder(r.Body).Decode(&list)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong request"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		updatedList, err := service.UpdateTodoList(id, list)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not create a Todo List"))
			return
		}

		json.NewEncoder(w).Encode(updatedList)
	}
}

func deleteTodoList(service services.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		slog.Info("DELETE /lists/" + idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong id format"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = service.DeleteTodoList(id)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not delete a Todo List: %v", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}


func createTodo(service services.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		slog.Info("DELETE /lists/" + idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong id format"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		slog.Info("POST /lists/" + idStr + "/todos")

		var todo models.Todo
		err = json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong request"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		createdTodo, err := service.CreateTodo(id, todo)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not create a Todo"))
			return
		}

		json.NewEncoder(w).Encode(createdTodo)
	}
}

func updateTodo(service services.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		listIdStr := r.PathValue("list_id")
		listId, err := strconv.Atoi(listIdStr)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong list id format"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		todoIdStr := r.PathValue("todo_id")
		todoId, err := strconv.Atoi(todoIdStr)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong todo id format"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}


		var todo models.Todo
		err = json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong request"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		updatedTodo, err := service.UpdateTodo(listId, todoId, todo)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not update a Todo"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(updatedTodo)
	}
}

func deleteTodo(service services.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todoIdStr := r.PathValue("todo_id")
		todoId, err := strconv.Atoi(todoIdStr)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong todo id format"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}


		err = service.DeleteTodo(todoId)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not delete a todo"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
