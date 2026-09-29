package main

import "fmt"

// Blank identifier (_) is used to ignore a value.
// USES 
// 1. To ignore a value returned by a function.4
// 2. To ignore the index of a slice or array when using a for loop.
// 3. TO REMOVE THE UNUSED VARIABLE ERROR
// 

func main() {

	// Declare two variables.
	a, b := 10, 20

	// Ignore the value of a.
	// Assign the value of b to c.
	_, c := a, b

	fmt.Println(c) // Output: 20

	// Assign the value of a to d.
	// Ignore the value of b.
	d, _ := a, b

	fmt.Println(d) // Output: 10

	// Here, we are not ignoring any value.
	// The value of a is assigned to e.
	// The value of b is assigned to f.
	e, f := a, b

	fmt.Println(e, f) // Output: 10 20

	// Here, we are also not ignoring any value.
	// The value of a is assigned to g.
	// The value of b is assigned to h.
	g, h := a, b

	fmt.Println(g, h) // Output: 10 20
}