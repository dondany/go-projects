package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/dondany/go-projects/to-do-rest/internal/models"
	"github.com/dondany/go-projects/to-do-rest/internal/services"
)

type TodoHandler struct {
	service services.TodoService
}

func NewTodoHandler(service services.TodoService) TodoHandler {
	return TodoHandler{
		service: service,
	}
}

func (h TodoHandler) GetTodoLists(w http.ResponseWriter, r *http.Request) {
	slog.Info("GET /lists")
	lists, err := h.service.GetTodoLists()
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not fetch lists"))
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(lists)
}

func (h TodoHandler) GetTodoList(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	slog.Info("GET /lists/" + idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Wrong id format"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	list, err := h.service.GetTodoList(id)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not fetch a list"))
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(list)
}

func (h TodoHandler) CreateTodoList(w http.ResponseWriter, r *http.Request) {
	slog.Info("POST /lists/")

	var list models.TodoList
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Wrong request"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdList, err := h.service.CreateTodoList(list)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not create a Todo List"))
		return
	}

	json.NewEncoder(w).Encode(createdList)
}

func (h TodoHandler) UpdateTodoList(w http.ResponseWriter, r *http.Request) {
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

	updatedList, err := h.service.UpdateTodoList(id, list)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not create a Todo List"))
		return
	}

	json.NewEncoder(w).Encode(updatedList)
}

func (h TodoHandler) DeleteTodoList(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	slog.Info("DELETE /lists/" + idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Wrong id format"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTodoList(id)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not delete a Todo List: %v", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
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

	createdTodo, err := h.service.CreateTodo(id, todo)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not create a Todo"))
		return
	}

	json.NewEncoder(w).Encode(createdTodo)
}

func (h TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
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

	updatedTodo, err := h.service.UpdateTodo(listId, todoId, todo)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not update a Todo"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedTodo)
}

func (h TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoIdStr := r.PathValue("todo_id")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Wrong todo id format"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTodo(todoId)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not delete a todo"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
