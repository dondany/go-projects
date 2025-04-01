export interface TodoList {
    ID: number,
    UserID: number,
    Name?: string,
    CreatedAt?: string 
    Todos?: Todo[]
}

export interface CreateTodoListRequest {
    UserID: number,
    Name: string
}

export interface DeleteTodoListRequest {
    UserID: number,
    ID: number
}

export interface Todo {
    ID: number,
    ListID: number,
    Name: string, 
    Completed: boolean,
    CretedAt: string
}

export interface CreateTodoRequest {
    UserID: number,
    ListID: number,
    Name: string,
}

export interface UpdateTodoRequest {
    ID: number,
    UserID: number,
    ListID: number,
    Name: string,
    Completed: boolean,
}

export interface DeleteTodoRequest {
    ID: number,
    UserID: number,
    ListID: number
}

