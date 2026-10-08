package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {

	dsn := "host=localhost port=5432 user=postgres dbname=go_crud sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	return db, nil
}