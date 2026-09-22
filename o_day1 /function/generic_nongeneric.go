package main

import "fmt"

func multiply(a, b int) int { // no generic function
	// here it only worked on the int data type, if we want to use it for other data types like float, string etc. then we have to create a new function for that data type.
	return a * b
}
func add[T int | float64](a T, b T) T { // generic function, it can work on multiple data types like int, float64 etc.    T is a type parameter.   yse to indicate the parameter

	// uses to handle multiple data types in a single function, it is a new feature in go 1.18 and above.

	return a + b

}

func first[T any](items []T) T {
	return items[2]
}

func main() {
	result := multiply(5, 10)
	fmt.Println(result)

	result2 := add(5.0, 10.0)
	fmt.Println(result2)

	numbers := []int{10, 20, 30}
	names := []string{"Raghav", "Ram", "Hari"}

	fmt.Println(first(numbers))
	fmt.Println(first(names))

}
