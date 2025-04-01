import { Route } from '@angular/router';

export const LISTS_ROUTES: Route[] = [
  {
    path: '',
    loadComponent: () => import('./lists.component'),
  },
  {
    path: ':id',
    loadComponent: () => import('./list/list.component'),
  },
];
