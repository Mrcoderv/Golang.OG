package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"

	"learning/2database/Using_sqlx/database"
	"learning/2database/Using_sqlx/student"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := runMenu(db, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func runMenu(db *sqlx.DB, input io.Reader, output io.Writer) error {
	reader := bufio.NewReader(input)
	if output == nil {
		output = io.Discard
	}

	for {
		printMenu(output)

		choice, err := readInt(reader, output, "Choose an option: ")
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			fmt.Fprintln(output, err)
			continue
		}

		switch choice {
		case 1:
			if err := student.CreateTable(db); err != nil {
				return fmt.Errorf("create students table: %w", err)
			}
			fmt.Fprintln(output, "Students table is ready.")
		case 2:
			name, err := readText(reader, output, "Student name: ")
			if err != nil {
				return err
			}
			id, err := student.CreateStudent(db, name)
			if err != nil {
				return fmt.Errorf("create student: %w", err)
			}
			fmt.Fprintf(output, "Student created with ID %d.\n", id)
		case 3:
			if err := printStudents(db, output, "Students:"); err != nil {
				return err
			}
		case 4:
			if err := updateStudent(db, reader, output); err != nil {
				return err
			}
		case 5:
			if err := deleteStudent(db, reader, output); err != nil {
				return err
			}
		case 6:
			fmt.Fprintln(output, "Exiting the application.")
			return nil

		default:
			fmt.Fprintln(output, "Invalid option.")
		}
	}
}

func printMenu(output io.Writer) {
	fmt.Fprintln(output, "\nStudent management")
	fmt.Fprintln(output, "1. Create students table")
	fmt.Fprintln(output, "2. Create a student")
	fmt.Fprintln(output, "3. Get all students")
	fmt.Fprintln(output, "4. Update a student")
	fmt.Fprintln(output, "5. Delete a student")
	fmt.Fprintln(output, "6. Exit")
}

func readInt(reader *bufio.Reader, output io.Writer, prompt string) (int, error) {
	line, err := readLine(reader, output, prompt)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(line)
	if err != nil {
		return 0, fmt.Errorf("enter a number")
	}
	return value, nil
}

func readText(reader *bufio.Reader, output io.Writer, prompt string) (string, error) {
	for {
		value, err := readLine(reader, output, prompt)
		if err != nil {
			return "", err
		}
		if value != "" {
			return value, nil
		}
		fmt.Fprintln(output, "Value cannot be empty.")
	}
}

func readLine(reader *bufio.Reader, output io.Writer, prompt string) (string, error) {
	fmt.Fprint(output, prompt)
	line, err := reader.ReadString('\n')
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return "", err
		}
		if strings.TrimSpace(line) == "" {
			return "", io.EOF
		}
	}
	return strings.TrimSpace(line), nil
}

func updateStudent(db *sqlx.DB, reader *bufio.Reader, output io.Writer) error {
	id, err := readInt(reader, output, "Student ID: ")
	if err != nil {
		return err
	}
	name, err := readText(reader, output, "New student name: ")
	if err != nil {
		return err
	}
	if err := student.UpdateStudent(db, id, name); err != nil {
		return fmt.Errorf("update student: %w", err)
	}
	return printStudents(db, output, "\nAfter update:")
}

func deleteStudent(db *sqlx.DB, reader *bufio.Reader, output io.Writer) error {
	id, err := readInt(reader, output, "Student ID: ")
	if err != nil {
		return err
	}
	if err := student.DeleteStudent(db, id); err != nil {
		return fmt.Errorf("delete student: %w", err)
	}
	return printStudents(db, output, "\nAfter delete:")
}

func printStudents(db *sqlx.DB, output io.Writer, title string) error {
	students, err := student.GetStudents(db)
	if err != nil {
		return fmt.Errorf("get students: %w", err)
	}

	fmt.Fprintln(output, title)
	for _, current := range students {
		fmt.Fprintf(output, "ID: %d, Name: %s\n", current.ID, current.Name)
	}
	return nil
}


