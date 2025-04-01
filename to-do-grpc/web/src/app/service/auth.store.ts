import { computed, inject } from '@angular/core';
import { Router } from '@angular/router';
import { LoginRequest, RegisterRequest, User } from '../model/user';
import { UserService } from './user.service';
import { patchState, signalStore, withComputed, withMethods, withState } from '@ngrx/signals';
import { rxMethod } from '@ngrx/signals/rxjs-interop';
import { pipe, switchMap, tap } from 'rxjs';

export interface AuthState {
  user: User | null;
}

export const AuthStore = signalStore(
  { providedIn: 'root' },
  withState<AuthState>({
    user: localStorage.getItem('user') ? JSON.parse(localStorage.getItem('user')!) : null
  }),
  withMethods(
    (store, userService = inject(UserService), router = inject(Router)) => {
        return {
            login: rxMethod<LoginRequest>(
                pipe(
                    switchMap(login => userService.loginUser(login)),
                    tap(response => patchState(store, { user: response})),
                    tap(response => localStorage.setItem('user', JSON.stringify(response))),
                    tap(() => router.navigate(['/lists'])),
                )
            ),
            register: rxMethod<RegisterRequest> (
              pipe(
                switchMap(register => userService.registerUser(register)),
                tap(response => patchState(store, { user: response})),
                tap(response => localStorage.setItem('user', JSON.stringify(response))),
                tap(() => router.navigate(['/lists'])),
              )
            )
        }
    }
  ),
  withComputed(store => ({
    
  })

  )
);
