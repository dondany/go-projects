package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/dondany/go-projects/to-do-grpc/pb"
	"github.com/dondany/go-projects/to-do-grpc/todo"
)

type TodoHandler struct {
	service pb.TodoServiceClient
}

func NewTodoHandler(service pb.TodoServiceClient) TodoHandler {
	return TodoHandler{
		service: service,
	}
}

func (h TodoHandler) GetTodoLists(w http.ResponseWriter, r *http.Request) {
	slog.Info(r.Method + " " + r.URL.Path)
	userIdStr := r.PathValue("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "Wrong id format"))
		return
	}

	response, err := h.service.GetTodoLists(r.Context(), &pb.TodoListFilter{
		UserId: int32(userId),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "Could not fetch lists"))
	}

	lists := make([]*todo.TodoList, len(response.Lists))
	for i, l := range response.Lists {
		list := todo.TodoList{
			ID: l.Id,
			Name: l.Name,
			UserID: l.UserId,
			CreatedAt: l.CreatedAt.AsTime(),
		}
		lists[i] = &list
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lists)
}

func (h TodoHandler) GetTodoList(w http.ResponseWriter, r *http.Request) {
	slog.Info(r.Method + " " + r.URL.Path)
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "Wrong id format"))
		return
	}

	response, err := h.service.GetTodoList(r.Context(), &pb.ID{Id: int32(id)})
	if err != nil {
		if strings.Contains(err.Error(), "PermissionDenied") {
			w.WriteHeader(http.StatusUnauthorized)
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	todos := make([]todo.Todo, len(response.Todos))
	for i, t := range response.Todos {
		todos[i] = todo.Todo{
			ID: t.Id,
			ListID: t.ListId,
			Name: t.Name,
			Completed: t.Completed,
			CreatedAt: t.CreatedAt.AsTime(),
		}
	}

	list := todo.TodoList{
		ID: response.Id,
		Name: response.Name,
		UserID: response.UserId,
		Todos: todos,
		CreatedAt: response.CreatedAt.AsTime(),
	}

	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not fetch a list"))
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func (h TodoHandler) CreateTodoList(w http.ResponseWriter, r *http.Request) {
	slog.Info(r.Method + " " + r.URL.Path)
	userIdStr := r.PathValue("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Wrong id format"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var list todo.TodoList
	err = json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Wrong request"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := &pb.TodoListRequest{
		Name: list.Name,
		UserId: int32(userId),
	}
	response, err := h.service.CreateTodoList(r.Context(), req)
	fmt.Println(err)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not create a Todo List"))
		return
	}

	createdList := todo.TodoList{
		ID: response.Id,
		Name: response.Name,
		CreatedAt: response.CreatedAt.AsTime(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdList)
}

func (h TodoHandler) UpdateTodoList(w http.ResponseWriter, r *http.Request) {
	slog.Info(r.Method + " " + r.URL.Path)
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Wrong id format"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var list todo.TodoList
	err = json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "Wrong request"))
		return
	}

	req := &pb.UpdateTodoListRequest{
		Id: int32(id),
		Name: list.Name,
	}

	res, err := h.service.UpdateTodoList(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "Could not update a Todo List"))
		return
	}

	updatedList := todo.TodoList{
		ID: res.Id,
		Name: res.Name,
		CreatedAt: res.CreatedAt.AsTime(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedList)
}

func (h TodoHandler) DeleteTodoList(w http.ResponseWriter, r *http.Request) {
	slog.Info(r.Method + " " + r.URL.Path)
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "Wrong id format"))
		return
	}

	_, err = h.service.DeleteTodoList(r.Context(), &pb.ID{Id: int32(id)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "Could not delete a Todo List: %v", err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	slog.Info(r.Method + " " + r.URL.Path)
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "Wrong id format"))
		return
	}

	var parsedTodo todo.Todo
	err = json.NewDecoder(r.Body).Decode(&parsedTodo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "Wrong request"))
		return
	}

	req := &pb.TodoRequest{
		ListId: int32(id),
		Name: parsedTodo.Name,
	}

	res, err := h.service.CreateTodo(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "Could not create a Todo"))
		return
	}

	createdTodo := todo.Todo{
		ID: res.Id,
		ListID: res.ListId,
		Name: res.Name,
		Completed: res.Completed,
		CreatedAt: res.CreatedAt.AsTime(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdTodo)
}

func (h TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	slog.Info(r.Method + " " + r.URL.Path)
	listIdStr := r.PathValue("list_id")
	_, err := strconv.Atoi(listIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "Wrong list id format"))
		return
	}

	todoIdStr := r.PathValue("todo_id")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "Wrong todo id format"))
		return
	}

	var parsedTodo todo.Todo
	err = json.NewDecoder(r.Body).Decode(&parsedTodo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "Wrong request"))
		return
	}

	req := &pb.TodoUpdateRequest{
		Id: int32(todoId),
		Name: parsedTodo.Name,
		Completed: parsedTodo.Completed,
	}

	res, err := h.service.UpdateTodo(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "Could not update a Todo"))
		return
	}

	updatedTodo := todo.Todo{
		ID: res.Id,
		ListID: res.ListId,
		Name: res.Name,
		Completed: res.Completed,
		CreatedAt: res.CreatedAt.AsTime(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTodo)
}

func (h TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	slog.Info(r.Method + " " + r.URL.Path)
	todoIdStr := r.PathValue("todo_id")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "Wrong todo id format"))
		return
	}

	_, err = h.service.DeleteTodo(r.Context(), &pb.ID{Id: int32(todoId)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "Could not delete a todo"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
