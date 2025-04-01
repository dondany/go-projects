import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import {
  FormBuilder,
  FormControl,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { LoginRequest, RegisterRequest } from '../../model/user';
import { AuthStore } from '../../service/auth.store';

@Component({
  standalone: true,
  selector: 'app-auth',
  template: `
    <div class="p-2 flex flex-col items-center bg-slate-300">
      @if(isSignIn) {
        <div class="flex flex-col items-center">
        <span>Sign In</span>
        <form [formGroup]="signInForm" (ngSubmit)="onSingInSubmit()">
          <div>
            <label for="email">Email</label>
            <input
              type="text"
              formControlName="email"
              class="p-2 mb-1 bg-white rounded cursor-pointer hover:bg-slate-100 flex justify-between"
            />
          </div>
          <div>
            <label for="password">Password</label>
            <input
              type="password"
              formControlName="password"
              class="p-2 mb-1 bg-white rounded cursor-pointer hover:bg-slate-100 flex justify-between"
            />
          </div>
          <button
            type="submit"
            class="p-2 bg-white rounded hover:cursor-pointer hover:bg-slate-200"
          >
            Sign In
          </button>
        </form>
        <div>
          Don't have an account? <a class="text-blue-500 hover:cursor-pointer" (click)="isSignIn=false">Sign up</a>
        </div>
      </div>
      } @else {
      <div class="flex flex-col items-center">
        <span>Sign Up</span>
        <form [formGroup]="signUpForm" (ngSubmit)="onSignUpSubmit()">
        <div>
            <label for="name">Name</label>
            <input
              type="text"
              formControlName="name"
              class="p-2 mb-1 bg-white rounded cursor-pointer hover:bg-slate-100 flex justify-between"
            />
          </div>
          <div>
            <label for="email">Email</label>
            <input
              type="text"
              formControlName="email"
              class="p-2 mb-1 bg-white rounded cursor-pointer hover:bg-slate-100 flex justify-between"
            />
          </div>
          <div>
            <label for="password">Password</label>
            <input
              type="password"
              formControlName="password"
              class="p-2 mb-1 bg-white rounded cursor-pointer hover:bg-slate-100 flex justify-between"
            />
          </div>
          <button
            type="submit"
            class="p-2 bg-white rounded hover:cursor-pointer hover:bg-slate-200"
          >
            Sign In
          </button>
        </form>
        <div>
          Already have an account? <a class="text-blue-500 hover:cursor-pointer" (click)="isSignIn=true">Sign In</a>
        </div>
      </div>
      }
    </div>
  `,
  imports: [ReactiveFormsModule],
})
export default class AuthComponent {
  authStore = inject(AuthStore);
  isSignIn = true;

  signInForm = inject(FormBuilder).group({
    email: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required),
  });

  signUpForm = inject(FormBuilder).group({
    name: new FormControl('', Validators.required),
    email: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required),
  });

  onSingInSubmit() {
    const signInFormValue = this.signInForm.getRawValue();

    const login: LoginRequest = {
      email: signInFormValue?.email || '',
      password: signInFormValue?.password || '',
    };

    this.authStore.login(login);
  }

  onSignUpSubmit() {
    const signUpFormValue = this.signUpForm.getRawValue();

    const signup: RegisterRequest = {
      name: signUpFormValue?.name || '',
      email: signUpFormValue?.email || '',
      password: signUpFormValue?.password || '',
    };
  

    this.authStore.register(signup)
  }
}
