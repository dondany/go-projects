import { HttpClient, HttpHeaders } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { CreateTodoListRequest, CreateTodoRequest, DeleteTodoListRequest, DeleteTodoRequest, Todo, TodoList, UpdateTodoRequest } from '../model/todo';
import { AuthStore } from './auth.store';


@Injectable({
  providedIn: 'root',
})
export class TodoService {
  http = inject(HttpClient)
  authStore = inject(AuthStore)

  getLists(userId: number) {
    const headers = this.createHeaders();
    return this.http.get<TodoList[]>(`/api/${userId}/lists`, { headers })
  }

  getList(userId: number, listId: number) {
    const headers = this.createHeaders();
    return this.http.get<TodoList>(`/api/${userId}/lists/${listId}`, { headers })
  }

  createTodoList(userId: number, list: CreateTodoListRequest) {
    const headers = this.createHeaders();
    return this.http.post<TodoList>(`/api/${userId}/lists`, { Name: list.Name }, { headers })
  }
  
  deleteTodoList(userId: number, list: DeleteTodoListRequest) {
    const headers = this.createHeaders();
    return this.http.delete<TodoList>(`/api/${userId}/lists/${list.ID}`, { headers })
  }

  createTodo(userId: number, listId: number, todo: CreateTodoRequest) {
    const headers = this.createHeaders();
    return this.http.post<Todo>(`/api/${userId}/lists/${listId}/todos`, { Name: todo.Name }, { headers })
  }

  updateTodo(userId: number, listId: number, todo: UpdateTodoRequest) {
    const headers = this.createHeaders();
    return this.http.put<Todo>(`/api/${userId}/lists/${listId}/todos/${todo.ID}`, { Name: todo.Name, Completed: todo.Completed}, { headers })
  }

  deleteTodo(userId: number, listId: number, todo: DeleteTodoRequest) {
    const headers = this.createHeaders();
    return this.http.delete(`/api/${userId}/lists/${listId}/todos/${todo.ID}`, { headers })
  }

  private createHeaders() {
    return new HttpHeaders({
      'Authorization': `Bearer ${this.authStore.user()!.Token}`
    });
  }
  
}