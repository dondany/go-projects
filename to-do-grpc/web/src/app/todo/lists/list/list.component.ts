import { CommonModule } from '@angular/common';
import { Component, inject, input, OnInit } from '@angular/core';
import { TodoService } from '../../../service/todo.service';
import { TodoStore } from '../../../service/todo.store';
import { CreateTodoRequest, DeleteTodoRequest, Todo, TodoList, UpdateTodoRequest } from '../../../model/todo';
import {
  FormBuilder,
  FormControl,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';

@Component({
  standalone: true,
  selector: 'app-list',
  template: `
    <div class="p-2 flex flex-col items-center bg-slate-300">
      <h2 class="p2 mb-2 text-lg">{{ todoStore.currentList()?.Name }}</h2>
      <ul class="w-full">
        @for (todo of todoStore.currentList()?.Todos; track $index) {
        <li
          class="p-2 mb-1 bg-white rounded cursor-pointer hover:bg-slate-100 flex justify-between"
          (click)="toggleTodo(todo)"
        >
          <div class="flex gap-3">
            <div class="size-4">
              @if (todo.Completed) {
              <span class="material-symbols-outlined"> check </span>
              } @else { }
            </div>
            <span>{{ todo.Name }}</span>
          </div>

          <button (click)="deleteTodo(todo); $event.stopPropagation()">
            <span class="material-symbols-outlined"> delete </span>
          </button>
        </li>
        }
      </ul>
      <div class="w-full">
        <form
          [formGroup]="newTodoForm"
          (ngSubmit)="onSubmit()"
          class="w-full flex gap-1"
        >
          <div class="w-full">
            <input
              type="text"
              formControlName="todo"
              class="w-full p-2 mb-1 bg-white rounded cursor-pointer hover:bg-slate-100 flex justify-between"
              placeholder="Add new todo..."
            />
          </div>
          <button
            type="submit"
            class="p-2 bg-white rounded hover:cursor-pointer hover:bg-slate-200"
          >
            Add
          </button>
        </form>
      </div>
    </div>
  `,
  imports: [CommonModule, ReactiveFormsModule],
})
export default class ListComponent implements OnInit {
  id = input.required<number>();
  todoStore = inject(TodoStore);

  newTodoForm = inject(FormBuilder).group({
    todo: new FormControl('', Validators.required),
  });

  ngOnInit(): void {
    this.todoStore.loadCurrentList({ UserID: 1, ID: this.id()})
  }

  onSubmit() {
    const newTodoFormValue = this.newTodoForm.getRawValue();
    if (!newTodoFormValue.todo) {
      return;
    }

    const todo: CreateTodoRequest = {
      ListID: this.id(),
      UserID: 1,
      Name: newTodoFormValue.todo,
    };

    this.todoStore.createTodo(todo);
    this.newTodoForm.reset();
  }

  toggleTodo(todo: Todo) {
    const updateTodo: UpdateTodoRequest = {
      ID: todo.ID,
      ListID: this.id(),
      UserID: 1,
      Name: todo.Name,
      Completed: !todo.Completed
    };

    this.todoStore.updateTodo(updateTodo)
  }

  deleteTodo(todo: Todo) {
    const deleteTodo: DeleteTodoRequest = {
      ID: todo.ID,
      ListID: this.id(),
      UserID: 1,
    };

    this.todoStore.deleteTodo(deleteTodo)
  }
}
