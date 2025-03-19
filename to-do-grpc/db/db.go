package db

import (
	"database/sql"
	"fmt"
)

func Connect() (*sql.DB, error) {
	dsn := "host=localhost port=5432 user=admin password=admin dbname=db sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open db connection")
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("could not ping the fb")
	}
	return db, nil
}
