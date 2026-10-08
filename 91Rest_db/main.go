package main

import (
	"log"
	"net/http"

	"learning/91Rest_db/database"
	"learning/91Rest_db/student"
)

func main() {

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
