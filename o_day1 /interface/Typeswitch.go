package main

import "fmt"

func checkType(value interface{}) {
	switch v := value.(type) { // it is used to determine the tpe of the value ;

	case int:
		fmt.Println("Integer:", v)

	case string:
		fmt.Println("String:", v)

	case bool:
		fmt.Println("Boolean:", v)

	case float64:
		fmt.Println("Float:", v)

	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	checkType(100)
	checkType("Hello")
	checkType(3.14)
	list := []int{1, 2, 3}
	checkType(list) // Unknown type
	checkType([]int{1, 2, 3}) // Unknown type'/ 
	checkType(struct{ Name string }{Name: "John"}) // Unknown type
}
