 /* 
package main

import "fmt"

// Defining a struct   to store the data of a person     
// struct is the collection of  data field to store on the single object
// Each field in the struct has a name and a type.
// strcut is defined on the outside of the main function and can be used in the main function

type persons struct {
    name string
    age  int
}

// Defining a method with struct receiver
func (p person) display() {
    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
}

func main() {
    // Creating an instance of the struct
    a := person{name: "a", age: 25}
    
    // Calling the method
    a.display()
}


package main

import "fmt"

// Defining a struct
type person struct {
    name string
}

// Method with pointer receiver to modify data
func (p *person) changeName(newName string) {
    p.name = newName
}

func main() {
    a := person{name: "a"}
    
    fmt.Println("Before:", a.name)
    
    // Calling the method to change the name
    a.changeName("b")
    
    fmt.Println("After:", a.name)
}
*/

/*Method                            	Function
Contains a receiver      	        Does not contain a receiver
Methods with the same name   	    Functions with the same name but 
defined in the program              different types are not allowed
b 
Cannot be used as a first-orde     Can be used as first-order objects
 object	
*/     
package main
import "fmt"

type person struct {
    name string
}
// Method with pointer receiver
func (p *person) updateName(newName string) {
    p.name = newName
	fmt.Println("Inside pointer method:", &p.name)

}

// Method with value receiver
func (p person) showName() {
    fmt.Println("Name:", p.name)
}

func main() {
    a := person{name: "a"}
    
    // Calling pointer method with value
    a.updateName("b")
    fmt.Println("After pointer method:", a.name)
    
    // Calling value method with pointer
    (&a).showName()
}