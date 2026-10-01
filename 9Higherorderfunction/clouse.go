//clouse.go

package main

import "fmt"

func main() {
	
	// example of closure
	add := func(a int) func(int) int {
		return func(b int) int {
			return a + b
		}
	}

	add5 := add(5)
	result := add5(3)
	fmt.Println("Result:", result) // Output: Result: 8
}