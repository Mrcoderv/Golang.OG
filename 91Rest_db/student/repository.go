package student

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// CreateTable creates the students table.
func CreateTable(db *sqlx.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS students (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL
		);
	`

	_, err := db.Exec(query)
	return err
}

// CreateStudent adds a new student.
func CreateStudent(db *sqlx.DB, name string) (Student, error) {
	var student Student

	query := `
		INSERT INTO students (name)
		VALUES ($1)
		RETURNING id, name;
	`

	err := db.Get(&student, query, name)
	fmt.Println("student added:", student)
	return student, err
}

// GetStudents returns all students.
func GetStudents(db *sqlx.DB) ([]Student, error) {
	var students []Student

	query := `
		SELECT id, name
		FROM students
		ORDER BY id;
	`

	err := db.Select(&students, query)
	fmt.Println("students retrieved:")

	return students, err
}

// GetStudent returns one student.
func GetStudent(db *sqlx.DB, id int) (Student, error) {
	var student Student

	query := `
		SELECT id, name
		FROM students
		WHERE id = $1;
	`

	err := db.Get(&student, query, id)

	return student, err
}

// UpdateStudent updates a student's name.
func UpdateStudent(db *sqlx.DB, id int, name string) (Student, error) {
	var student Student

	query := `
		UPDATE students
		SET name = $1
		WHERE id = $2
		RETURNING id, name;
	`

	err := db.Get(&student, query, name, id)

	return student, err
}

// DeleteStudent deletes a student.
func DeleteStudent(db *sqlx.DB, id int) error {
	query := `
		DELETE FROM students
		WHERE id = $1;
	`

	_, err := db.Exec(query, id)

	return err
}
