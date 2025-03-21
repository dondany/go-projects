package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := "disable"

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	host,
	port,
	user,
	password,
	dbname,
	sslmode)

	slog.Info("opening connection to db")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	slog.Info("succesfuly opened connection to db")

	slog.Info("pinging db")
	if err = db.Ping(); err != nil {
		db.Close()
		panic(err)
	}
	slog.Info("successfuly pinged db")
	return db, nil
}
