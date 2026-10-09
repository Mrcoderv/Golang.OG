package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open(
		"postgres",
		"postgresql://postgres://mrrv:postgrace@localhost:5432/go_crud?sslmode=disable",
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")



	// creating the db
	err = db.QueryRow("CREATE TABLE IF NOT EXISTS Teacher (id SERIAL PRIMARY KEY, name TEXT, email TEXT)").Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully.")




}
func addTeacher(name string, email string) {
	db, err := sql.Open(
		"postgres",
		"postgresql://postgres:postgres@