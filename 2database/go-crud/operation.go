package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type SQLUser struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func CreateUserInteractive(db *sql.DB) {
	var name, email string
	var age int

	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	fmt.Print("Enter email: ")
	fmt.Scan(&email)

	fmt.Print("Enter age: ")
	fmt.Scan(&age)

	var id int

	err := db.QueryRow(
		`INSERT INTO users (name, email, age)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		name,
		email,
		age,
	).Scan(&id)

	if err != nil {
		log.Println("Error:", err)
		return
	}

	fmt.Println("User created with ID:", id)
}

func GetUser(db *sql.DB) {
	var id int

	fmt.Print("Enter user ID: ")
	fmt.Scan(&id)

	var user SQLUser

	err := db.QueryRow(
		`SELECT id, name, email, age
		 FROM users
		 WHERE id = $1`,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Age,
	)

	if err != nil {
		log.Println("Error:", err)
		return
	}

	fmt.Printf("ID: %d\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Age: %d\n", user.Age)
}

func UpdateUser(db *sql.DB) {
	var id int
	var name, email string
	var age int

	fmt.Print("Enter user ID: ")
	fmt.Scan(&id)

	fmt.Print("Enter new name: ")
	fmt.Scan(&name)

	fmt.Print("Enter new email: ")
	fmt.Scan(&email)

	fmt.Print("Enter new age: ")
	fmt.Scan(&age)

	_, err := db.Exec(
		`UPDATE users
		 SET name = $1, email = $2, age = $3
		 WHERE id = $4`,
		name,
		email,
		age,
		id,
	)

	if err != nil {
		log.Println("Error:", err)
		return
	}

	fmt.Println("User updated successfully.")
}

func DeleteUser(db *sql.DB) {
	var id int

	fmt.Print("Enter user ID: ")
	fmt.Scan(&id)

	_, err := db.Exec(
		`DELETE FROM users
		 WHERE id = $1`,
		id,
	)

	if err != nil {
		log.Println("Error:", err)
		return
	}

	fmt.Println("User deleted successfully.")
}

func GetAllUsers(db *sql.DB) {
	rows, err := db.Query(
		`SELECT id, name, email, age
		 FROM users
		 ORDER BY id`,
	)

	if err != nil {
		log.Println("Error:", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user SQLUser

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Age,
		)

		if err != nil {
			log.Println("Error:", err)
			return
		}

		fmt.Printf(
			"ID: %d | Name: %s | Email: %s | Age: %d\n",
			user.ID,
			user.Name,
			user.Email,
			user.Age,
		)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error:", err)
	}
}

func DeleteAllUsers(db *sql.DB) {
	_, err := db.Exec(`DELETE FROM users`)

	if err != nil {
		log.Println("Error:", err)
		return
	}

	fmt.Println("All users deleted successfully.")
}
