// package main

// import "fmt"
// // uses 1. 
// func main() {
// 	var value any = 100  // here the randomy value is assigned to the empty interface variable. The empty interface can hold any type of value.

// 	// Type assertion
// 	number, ok := value.(int)  // here the type assertion is used to check if the value stored in the empty interface variable is of type int. The result of the type assertion is stored in the variable 'number' and a boolean value 'ok' indicating whether the assertion was successful or not.

// 	if ok {
// 		fmt.Println("Value:", number)
// 		fmt.Println("Type assertion successful")
// 	} else {
// 		fmt.Println("Type assertion failed")

// 	}

// }


// example of type assertion on the injterface
package main

import "fmt"

// Interface
type Animal interface {
	Speak()
}

// Struct
type Dog struct {
	Name string
	Age  int
}

// Dog implements Animal
func (Dog) Speak() {
	fmt.Println("Woof!")
}

func main() {

	// Interface variable containing a Dog
	var animal Animal = Dog{
		Name: "Buddy" ,
		Age:  3,
	}

	// Type assertion
	dog, ok := animal.(Dog)

	if ok {
		fmt.Println("Name:", dog.Name)
		fmt.Println("Age:", dog.Age)
		dog.Speak()
	} else {
		fmt.Println("The value is not a Dog")
	}
}

