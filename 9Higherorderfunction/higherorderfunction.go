// higher orderfunction in go
// package main

// import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

// func calculate(a, b int, operation func(int, int) int) int {
// 	return operation(a, b)
// }

// func main() {
// 	result := calculate(10, 20, add)

// 	fmt.Println(result)
// }

// now lets keep seperate
// // taking the function as a pareameter
// package main

// import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

// func calculate(a, b int, operation func(int, int) int) int {
// 	return operation(a, b)
// }

// func main() {
// 	result := calculate(10, 20, add)

// 	fmt.Println(result)
// }
// operation key word is used to pass the function as the parameter to the calcualation.

// now returning the function from the function.
// func main() {
// 	innerFunc := outer()
// 	innerFunc() 
// }
// func outer() func() {
// 	x := 10

// 	return func() { //  we can directly return the function.
// 		fmt.Println(x)
// 	}
// }
