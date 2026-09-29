package main

import "fmt"

// printValue is a function that takes an empty interface as a parameter.
// The empty interface can be used to accept any type of value as the parameter.
func printValue(value interface{}) {    //. here the empty interface is used in as the parameter of the function.  The empty interface can be used to accept any type of value as the parameter.
	fmt.Println(value) 
}

func main() {
	printValue(100)
	printValue("Hello")
	printValue(true)
	printValue(3.14)
}