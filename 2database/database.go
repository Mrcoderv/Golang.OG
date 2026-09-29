// package main

// import (
// 	"fmt"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type Student struct {
// 	ID    int
// 	Name  string
// 	Email string
// }

// var db *gorm.DB

// func initDB() {
// 	var err error
// 	db, err = gorm.Open(postgres.Open("postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable"), &gorm.Config{})

// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Database connected.")

// }
// 	func addStudent(name string, email string) {
// 	student := Student{Name: name, Email: email}
// 	result := db.Create(&student)

// 	fmt.Printf("Student %s added successfully.\n", name)
// }

// func createStudentTable() {
// 	if err := db.AutoMigrate(&Student{}); err != nil {
// 		panic(err)
// 	}
// }
// func getStudent(id int) {
// 	var student Student
// 	result := db.First(&student, id)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	fmt.Printf("Student ID: %d, Name: %s, Email: %s\n", student.ID, student.Name, student.Email)
// }
// func removeStudent(id int) {
// 	var student Student
// 	result := db.First(&student, id)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	db.Delete(&student)
// 	fmt.Printf("Student with ID %d removed successfully.\n", id)

// }
// func updateStudent(id int, name string, email string) {
// 	var student Student
// 	result := db.First(&student, id)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	student.Name = name
// 	student.Email = email
// 	db.Save(&student)
// 	fmt.Printf("Student with ID %d updated successfully.\n", id)
// }

// func main() {
// 	fmt.Println("Starting the application...")
// 	initDB()
// 	createStudentTable()

// 	addStudent("rAGHAV", "raghav@example.com")
// addStudent("AGHAV", "Aghav@example.com")

// 	getStudent(1)

// 	updateStudent(1, "Sagar", "Sagar@exle.com")

// removeStudent(1)

// }
