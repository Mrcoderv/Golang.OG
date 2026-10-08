package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func EnsureSchema(ctx context.Context, db *pgxpool.Pool) error {
	_, err := db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(150) UNIQUE NOT NULL,
			age INT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)

	return err
}

func CreateUser(
	ctx context.Context,
	db *pgxpool.Pool,
	name string,
	email string,
	age int,
) (int, error) {

	var id int

	err := db.QueryRow(
		ctx,
		`
		INSERT INTO users (name, email, age)
		VALUES ($1, $2, $3)
		RETURNING id
		`,
		name,
		email,
		age,
	).Scan(&id)

	return id, err
}


func GetUserByID(
	ctx context.Context,
	db *pgxpool.Pool,
	id int,
) (*User, error)
func main() {
	ctx := context.Background()

	db, err := pgxpool.New(
		ctx,
		"postgres://postgres:postgres@localhost:5432/go_crud",
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL!")

	if err := EnsureSchema(ctx, db); err != nil {
		log.Fatal(err)
	}

	id, err := CreateUser(ctx, db, "John Doe", "john.doe@example.com", 30)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User created successfully with ID %d.\n", id)
}
