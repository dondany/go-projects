package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

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
	slog.Info("GET /lists")

	response, err := h.service.GetTodoLists(context.Background(), &pb.TodoListFilter{})
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not fetch lists"))
		w.WriteHeader(http.StatusInternalServerError)
	}

	lists := make([]*todo.TodoList, len(response.Lists))
	for i, l := range response.Lists {
		list := todo.TodoList{
			ID: l.Id,
			Name: l.Name,
			CreatedAt: l.CreatedAt.AsTime(),
		}
		lists[i] = &list
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

	response, err := h.service.GetTodoList(context.Background(), &pb.ID{Id: int32(id)})

	list := todo.TodoList{
		ID: response.Id,
		Name: response.Name,
		CreatedAt: response.CreatedAt.AsTime(),
	}

	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not fetch a list"))
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(list)
}

func (h TodoHandler) CreateTodoList(w http.ResponseWriter, r *http.Request) {
	slog.Info("POST /lists/")

	var list todo.TodoList
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Wrong request"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := &pb.TodoListRequest{
		Name: list.Name,
	}
	response, err := h.service.CreateTodoList(context.Background(), req)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not create a Todo List"))
		return
	}

	createdList := todo.TodoList{
		ID: response.Id,
		Name: response.Name,
		CreatedAt: response.CreatedAt.AsTime(),
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

	var list todo.TodoList
	err = json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Wrong request"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := &pb.UpdateTodoListRequest{
		Id: int32(id),
		Name: list.Name,
	}

	res, err := h.service.UpdateTodoList(context.Background(), req)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not create a Todo List"))
		return
	}

	updatedList := todo.TodoList{
		ID: res.Id,
		Name: res.Name,
		CreatedAt: res.CreatedAt.AsTime(),
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

	_, err = h.service.DeleteTodoList(context.Background(), &pb.ID{Id: int32(id)})
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

	var parsedTodo todo.Todo
	err = json.NewDecoder(r.Body).Decode(&parsedTodo)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Wrong request"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := &pb.TodoRequest{
		ListId: int32(id),
		Name: parsedTodo.Name,
	}

	res, err := h.service.CreateTodo(context.Background(), req)
	if err != nil {
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

	json.NewEncoder(w).Encode(createdTodo)
}

func (h TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	listIdStr := r.PathValue("list_id")
	_, err := strconv.Atoi(listIdStr)
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
	slog.Info("PUT /lists/" + listIdStr + "/todos/" + todoIdStr)

	var parsedTodo todo.Todo
	err = json.NewDecoder(r.Body).Decode(&parsedTodo)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Wrong request"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := &pb.TodoUpdateRequest{
		Id: int32(todoId),
		Name: parsedTodo.Name,
		Completed: parsedTodo.Completed,
	}

	res, err := h.service.UpdateTodo(context.Background(), req)
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not update a Todo"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedTodo := todo.Todo{
		ID: res.Id,
		ListID: res.ListId,
		Name: res.Name,
		Completed: res.Completed,
		CreatedAt: res.CreatedAt.AsTime(),
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

	_, err = h.service.DeleteTodo(context.Background(), &pb.ID{Id: int32(todoId)})
	if err != nil {
		w.Write(fmt.Appendf(nil, "Could not delete a todo"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
