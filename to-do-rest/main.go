package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/dondany/go-projects/to-do-rest/models"
	"github.com/dondany/go-projects/to-do-rest/repositories"
	"github.com/dondany/go-projects/to-do-rest/services.go"
)

func main() {
	mux := http.NewServeMux()

	repo := repositories.NewTodoMemoryRepository()
	service := services.NewTodoService(repo)

	mux.HandleFunc("GET /lists", getTodoLists(service))
	mux.HandleFunc("GET /lists/{name}", getTodoList(service))
	mux.HandleFunc("POST /lists", createTodoList(service))
	mux.HandleFunc("PATCH /lists/{name}", updateTodoList(service))
	mux.HandleFunc("DELETE /lists/{name}", deleteTodoList(service))

	mux.HandleFunc("POST /lists/{name}/todos", createTodo(service))
	mux.HandleFunc("PUT /lists/{name}/todos/{id}", updateTodo(service))
	mux.HandleFunc("DELETE /lists/{name}/todos/{id}", deleteTodo(service))

	err := http.ListenAndServe(":8081", mux)
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
		name := r.PathValue("name")
		slog.Info("POST /lists/" + name)

		list, err := service.GetTodoList(name)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not fetch a list"))
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(list)
	}
}

func createTodo(service services.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		slog.Info("POST /lists/" + name + "/todos")

		var todo models.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong request"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		createdTodo, err := service.CreateTodo(name, todo)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not create a Todo"))
			return
		}

		json.NewEncoder(w).Encode(createdTodo)
	}
}

func updateTodo(service services.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		idStr := r.PathValue("id")
		slog.Info("POST /lists/" + name + "/todos/" + idStr)
		id, err := strconv.Atoi(idStr)
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

		updatedTodo, err := service.UpdateTodo(name, id, todo)
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
		name := r.PathValue("name")
		idStr := r.PathValue("id")
		slog.Info("POST /lists/" + name + "/todos/" + idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong todo id format"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = service.DeleteTodo(name, id)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not delete a todo"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func createTodoList(service services.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		slog.Info("POST /lists/" + name + "/todos")

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
		name := r.PathValue("name")
		slog.Info("POST /lists/" + name + "/todos")

		var list models.TodoList
		err := json.NewDecoder(r.Body).Decode(&list)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Wrong request"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		updatedList, err := service.UpdateTodoList(name, list)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not create a Todo List"))
			return
		}

		json.NewEncoder(w).Encode(updatedList)
	}
}

func deleteTodoList(service services.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		slog.Info("POST /lists/" + name)

		err := service.DeleteTodoList(name)
		if err != nil {
			w.Write(fmt.Appendf(nil, "Could not delete a Todo List"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}