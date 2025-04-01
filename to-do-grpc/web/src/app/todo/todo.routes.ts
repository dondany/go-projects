import { Route } from '@angular/router';

export const TODO_ROUTES: Route[] = [
  {
    path: '',
    loadComponent: () => import('./home/home.component'),
  },
  {
    path: 'auth',
    loadComponent: () => import('./auth/auth.component'),
  },
  {
    path: 'lists',
    loadChildren: () =>
      import('./lists/lists.routes').then(m => m.LISTS_ROUTES),
  }
];
