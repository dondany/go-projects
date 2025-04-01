import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { LoginRequest, User } from '../model/user';
import { UserService } from './user.service';
import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { rxMethod } from '@ngrx/signals/rxjs-interop';
import { pipe, switchMap, tap, throwError } from 'rxjs';
import {
  CreateTodoListRequest,
  CreateTodoRequest,
  DeleteTodoListRequest,
  DeleteTodoRequest,
  TodoList,
  UpdateTodoRequest,
} from '../model/todo';
import { TodoService } from './todo.service';
import { AuthStore } from './auth.store';

export interface TodoState {
  lists: TodoList[];
  currentList: TodoList | null;
}

export const TodoStore = signalStore(
  { providedIn: 'root' },
  withState<TodoState>({
    lists: [],
    currentList: null,
  }),
  withMethods(
    (store, todoService = inject(TodoService), authStore = inject(AuthStore), router = inject(Router)) => {
      return {
        loadLists: rxMethod<void>(
          pipe(
            switchMap(() => todoService.getLists(authStore.user()!.ID)),
            tap((response) => patchState(store, { lists: response })),
            tap(() => router.navigate(['/lists']))
          )
        ),
        loadCurrentList: rxMethod<TodoList>(
          pipe(
            switchMap((list) => todoService.getList(authStore.user()!.ID, list.ID)),
            tap((response) => patchState(store, { currentList: response })),
            tap((response) => router.navigate(['/lists', response.ID]))
          )
        ),
        createTodoList: rxMethod<CreateTodoListRequest>(
          pipe(
            switchMap((list) => todoService.createTodoList(authStore.user()!.ID, list)),
            tap((r) => console.log('list', r)),
            switchMap(() => todoService.getLists(authStore.user()!.ID)),
            tap((response) => patchState(store, { lists: response })),
          )
        ),
        deleteTodoList: rxMethod<DeleteTodoListRequest>(
          pipe(
            switchMap((list) => todoService.deleteTodoList(authStore.user()!.ID, list)),
            switchMap(() => todoService.getLists(authStore.user()!.ID)),
            tap((response) => patchState(store, { lists: response })),
          )
        ),
        createTodo: rxMethod<CreateTodoRequest>(
          pipe(
            switchMap((todo) =>
              todoService.createTodo(authStore.user()!.ID, todo.ListID, todo)
            ),
            switchMap((todo) => todoService.getList(authStore.user()!.ID, todo.ListID)),
            tap((response) => patchState(store, { currentList: response }))
          )
        ),
        updateTodo: rxMethod<UpdateTodoRequest>(
          pipe(
            switchMap((todo) =>
              todoService.updateTodo(authStore.user()!.ID, todo.ListID, todo)
            ),
            switchMap((todo) => todoService.getList(authStore.user()!.ID, todo.ListID)),
            tap((response) => patchState(store, { currentList: response }))
          )
        ),
        deleteTodo: rxMethod<DeleteTodoRequest>(
          pipe(
            switchMap((todo) =>
              todoService.deleteTodo(authStore.user()!.ID, todo.ListID, todo)
            ),
            switchMap(() => todoService.getList(authStore.user()!.ID, store.currentList()!.ID)),
            tap((response) => patchState(store, { currentList: response }))
          )
        ),
      };
    }
  )
);
