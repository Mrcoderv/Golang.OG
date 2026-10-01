// without higher order functions
// package main 

// import "fmt"

// func add(a, b int) int {
// 	return a+ b
// }
// func subtract(a, b int) int {
// 	return a - b
// }
// func multiply(a, b int) int {
// 	return a * b
// }
// func divide(a, b int) int {
// 	return a / b
// }

// func main() {

// 	// result := add(10, 20)
// 	// result = subtract(10, 20)	
// 	// result = multiply(10, 20)
// 	// result = divide(10, 20)

// 	fmt.Println(result)
// }		


// with higher order functions

// package main
//  func add(a, b int) int {
//  	return a + b
//  }
//  func subtract(a, b int) int {
//  	return a - b
//  }
//  func multiply(a, b int) int {
//  	return a * b
//  }
//  func divide(a, b int) int {
//  	return a / b
//  }

//  func calculate(a, b int, operation func(int, int) int) int {
//  	return operation(a, b)
//  }

//  func main() {

//  	result := calculate(10, 20, add)   // just the parameter changes.
//  	// result = calculate(10, 20, subtract)
//  	// result = calculate(10, 20, multiply)
//  	// result = calculate(10, 20, divide)
//  	fmt.Println(result)
//  }