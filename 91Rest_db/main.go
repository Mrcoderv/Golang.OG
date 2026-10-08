package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"91Rest_db/database"
	"91Rest_db/student"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatalf("failed to load .env: %v", err)
	}

	// Connect to PostgreSQL
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("connected to PostgreSQL")

	err = student.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("table ready")

	studentHandler := &student.Handler{
		DB: db,
	}

	http.HandleFunc("/students", studentHandler.HandleStudents)
	http.HandleFunc("/students/", studentHandler.HandleStudent)

	log.Println("Server running on http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
