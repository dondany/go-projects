export interface User {
    ID: number;
    Name: string;
    Email: string;
    Token: string;
}

export interface RegisterRequest {
    name: string;
    email: string;
    password: string;
}

export interface LoginRequest {
    email: string;
    password: string;
}