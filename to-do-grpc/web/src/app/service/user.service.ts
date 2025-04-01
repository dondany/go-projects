import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { LoginRequest, RegisterRequest, User } from '../model/user';


@Injectable({
  providedIn: 'root',
})
export class UserService {
  http = inject(HttpClient)

  loginUser(login: LoginRequest) {
    return this.http.post<User>(`/api/user/login`, login);
  }

  registerUser(register: RegisterRequest) {
    return this.http.post<User>(`/api/user/register`, register);
  }
}