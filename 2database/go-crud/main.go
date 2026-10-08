package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
) (*User, error) {
	var user User

	err := db.QueryRow(ctx, `
		SELECT id, name, email, age
		FROM users
		WHERE id = $1
	`, id).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func main() {
	ctx := context.Background()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set; use postgres://USER:PASSWORD@localhost:5432/go_crud?sslmode=disable")
	}

	db, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
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
