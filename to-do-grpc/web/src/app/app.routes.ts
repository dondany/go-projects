import { Routes } from '@angular/router';

export const routes: Routes = [
    {
        path: '',
        loadChildren: () =>
          import('./todo/todo.routes').then(m => m.TODO_ROUTES),
      }
];
