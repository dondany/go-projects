import { CommonModule } from '@angular/common';
import { Component, inject, OnInit } from '@angular/core';
import { Router, RouterOutlet } from '@angular/router';
import { AuthStore } from '../../service/auth.store';
import { TodoStore } from '../../service/todo.store';
import { CreateTodoListRequest, DeleteTodoListRequest, TodoList } from '../../model/todo';
import { FormBuilder, FormControl, ReactiveFormsModule, Validators } from '@angular/forms';

@Component({
  standalone: true,
  selector: 'app-lists',
  template: `
  <div class="w-full p-2 rounded bg-slate-300 flex flex-col items-center">
  <h2 class="p2 mb-2 text-lg">My Todo Lists</h2>
  <ul class="w-full">
    @for (list of todoStore.lists(); track $index) {
    <li (click)="openList(list)" class="p-2 mb-1 bg-white rounded cursor-pointer hover:bg-slate-100 flex justify-between">
      <span>{{ list.Name }}</span>

      <button (click)="deleteTodoList(list); $event.stopPropagation()">
            <span class="material-symbols-outlined"> delete </span>
          </button>
    </li>

    }
  </ul>
  <div class="w-full">
        <form
          [formGroup]="newListForm"
          (ngSubmit)="onSubmit()"
          class="w-full flex gap-1"
        >
          <div class="w-full">
            <input
              type="text"
              formControlName="name"
              class="w-full p-2 mb-1 bg-white rounded cursor-pointer hover:bg-slate-100 flex justify-between"
              placeholder="Add new list..."
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
  <router-outlet />
</div>`,
  imports: [CommonModule, RouterOutlet, ReactiveFormsModule],
})
export default class ListsComponent implements OnInit {
  authStore = inject(AuthStore);
  todoStore = inject(TodoStore);

  newListForm = inject(FormBuilder).group({
    name: new FormControl('', Validators.required),
  });

  ngOnInit(): void {
    this.todoStore.loadLists();
  }

  openList(list: TodoList) {
    this.todoStore.loadCurrentList(list)
  }

  onSubmit() {
    const newListFormValue = this.newListForm.getRawValue();
        if (!newListFormValue.name) {
          return;
        }
    
        const list: CreateTodoListRequest = {
          UserID: 1,
          Name: newListFormValue.name,
        };
    
        this.todoStore.createTodoList(list);
        this.newListForm.reset();
  }

  deleteTodoList(list: TodoList) {
    const toBeDeleted: DeleteTodoListRequest = {
      ID: list.ID,
      UserID: list.UserID
    }

    this.todoStore.deleteTodoList(toBeDeleted)
  }
}
