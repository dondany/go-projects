package main

import (
	"database/sql"
	"net/http"

	"github.com/dondany/go-projects/to-do-rest/internal/handlers"
	"github.com/dondany/go-projects/to-do-rest/internal/repositories"
	"github.com/dondany/go-projects/to-do-rest/internal/services"
	_ "github.com/lib/pq"
)

func main() {
	db := connectDb()
	psqlRepo := repositories.NewTodoPostgreSqlRepository(db)
	service := services.NewTodoService(psqlRepo)
	handler := handlers.NewTodoHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /lists", handler.GetTodoLists)
	mux.HandleFunc("GET /lists/{id}", handler.GetTodoList)
	mux.HandleFunc("POST /lists", handler.CreateTodoList)
	mux.HandleFunc("PATCH /lists/{id}", handler.UpdateTodoList)
	mux.HandleFunc("DELETE /lists/{id}", handler.DeleteTodoList)
	mux.HandleFunc("POST /lists/{id}/todos", handler.CreateTodo)
	mux.HandleFunc("PUT /lists/{list_id}/todos/{todo_id}", handler.UpdateTodo)
	mux.HandleFunc("DELETE /lists/{list_id}/todos/{todo_giid}", handler.DeleteTodo)

	http.ListenAndServe(":8081", mux)
}

func connectDb() *sql.DB {
	dsn := "host=localhost port=5432 user=admin password=admin dbname=db sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		panic(err)
	}
	return db
}
