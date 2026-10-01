// package main

// import "fmt"

// type addingcontent interface { 
// 	int | string
// }

// func add[T addingcontent](a, b T) T {
// 	return a + b
// }

// func subtract(a, b int) int {
// 	return a - b
// }

// func calculate[T addingcontent](a, b T, operation func(T, T) T) T {
// 	return operation(a, b)
// }

// func main() {
// 	result := calculate(10, 20, add)
// 	result1 := calculate("Hello ", "World", add)
// 	result2 := subtract(10, 20)
// 	fmt.Println(result2)
// 	fmt.Println(result)
// 	fmt.Println(result1)
// }
