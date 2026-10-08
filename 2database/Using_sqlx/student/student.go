package student

import (
	"github.com/jmoiron/sqlx"
)

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

// CreateStudent adds a new student.
func CreateStudent(db *sqlx.DB, name string) error {
	query := `
		INSERT INTO students (name)
		VALUES ($1);
	`

	_, err := db.Exec(query, name)
	return err
}

// GetStudents gets all students.
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

// GetStudent gets one student by ID.
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

	_, err := db.Exec(query, name, id)
	return err
}

func DeleteStudent(db *sqlx.DB, id int) error {
	query := `
		DELETE FROM students
		WHERE id = $1;
	`

	_, err := db.Exec(query, id)
	return err
}
