// // imagining the three student doing the homework
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func task(channel chan string, studentName string) {
// 	time.Sleep(3 * time.Second) // Simulate a long-running task where the multiple go routines are working on the same task
// 	fmt.Println(studentName, "working ")
// 	channel <- studentName + " finished the task"
// }

// func main() {
// 	channel := make(chan string)

// 	defer close(channel)

// 	go task(channel, "Student_A")
// 	go task(channel, "Student_B")
// 	go task(channel, "Student_C")

// 	for result := range channel {
// 		fmt.Println(result)

// 	}

// 	// Wait for tasks to complete or be canceled
// }
