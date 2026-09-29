package main

import "fmt"

// USES
// TO CLOSET THE FILE , DATA BASE ,
//defer schedules a function call to execute when the surrounding function returns.
// 1. To close files, database connections, or other resources.

func main() {

	defer fmt.Println("This runs later") //  HERO KI ENTRY LAST MAI HOTI HAI

	fmt.Println("Hello")
	fmt.Println("Welcome")

	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third") // FIRST IN LAST OUT  // LAST IN FIRST OUT

	defer func() { // defer using anonymous function
		fmt.Println("Cleaning up...")
	}()
}

/*
defer works like a stack:

defer First
defer Second
defer Third

        ↓
Third     ← first
Second    ← second
First     ← last
*/
