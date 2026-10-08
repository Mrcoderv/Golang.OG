package main

import (
	"fmt"
	"log"

	"learning/2database/Using_sqlx/database"
	"learning/2database/Using_sqlx/student"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create students table
	err = student.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}

	// CREATE
	err = student.CreateStudent(db, "John")
	if err != nil {
		log.Fatal(err)
	}

	err = student.CreateStudent(db, "Alice")
	if err != nil {
		log.Fatal(err)
	}

	// READ
	students, err := student.GetStudents(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Students:")
	for _, s := range students {
		fmt.Printf("ID: %d, Name: %s\n", s.ID, s.Name)
	}

	err = student.UpdateStudent(db, 1, "John Updated")
	if err != nil {
		log.Fatal(err)
	}
	err = student.DeleteStudent(db, 2)
	if err != nil {
		log.Fatal(err)
	}
}
