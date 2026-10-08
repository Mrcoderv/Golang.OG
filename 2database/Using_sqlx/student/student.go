package student

import (
	"database/sql"
	"errors"

	"fmt"

	"github.com/jmoiron/sqlx"
)

// ErrNotFound is returned when an update or delete matches no row.
var ErrNotFound = errors.New("student not found")

type Student struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

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

// CreateStudent adds a new student and returns the generated ID.
func CreateStudent(db *sqlx.DB, name string) (int, error) {
	query := `
		INSERT INTO students (name)
		VALUES ($1)
		RETURNING id;
	`

	var id int
	err := db.Get(&id, query, name)
	return id, err
}

// all students.
func GetStudents(db *sqlx.DB) ([]Student, error) {
	var students []Student

	query := `
		SELECT id, name
		FROM students
		ORDER BY id;
	`

	err := db.Select(&students, query)
	return students, err
}

// student by ID.
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

// UpdateStudent changes a student's name.
func UpdateStudent(db *sqlx.DB, id int, name string) error {
	query := `
		UPDATE students
		SET name = $1
		WHERE id = $2;
	`

	res, err := db.Exec(query, name, id)
	if err != nil {
		return err
	}
	return requireRow(res, id)
}

// DeleteStudent removes a student by ID.
func DeleteStudent(db *sqlx.DB, id int) error {
	query := `
		DELETE FROM students
		WHERE id = $1;
	`

	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return requireRow(res, id)
}

// requireRow turns "0 rows affected" into ErrNotFound.
func requireRow(res sql.Result, id int) error {
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("id %d: %w", id, ErrNotFound)
	}
	return nil
}
