package main

import "fmt"

type speaker struct {    //used to store data
	name string
}
type Speaker interface {  // use to do task
	Speak()  // function signature   to define the behavior
	call()
}

func (s speaker) Speak() {
	fmt.Println("Hello, I am", s.name)
}

func (s speaker) call() {
	fmt.Println("Calling", s.name)
}

func main() {
	s := speaker{name: "Raghav"}  // here the instance is cerated of the struct and the value is assigned to the field of the struct.
	var sp Speaker = s   // here the instance of the struct is assigned to the interface variable.  The interface variable can hold any value that implements the interface.
	sp.Speak() // calling the method of the interface. 
	sp.call()
}
