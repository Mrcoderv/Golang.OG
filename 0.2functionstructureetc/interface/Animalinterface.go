// the main task of the
// interface is to define a contract that structs can implement.
// This allows for polymorphism, where different structs can be treated as the same type if they implement the same interface.

package main

import "fmt"

type Dog struct { // here the dog struct is created for strong the name of the dog
	Name string
}

type Cat struct {
	Name string
}
type Speaker interface { // here the interface is created which is used to define the behavior of the struct.  The interface can be implemented by any struct that has the same method signature.
	Speak()
	Move()
}

func (d Dog) Speak() { //  function is created which is used to define the behavior of the dog struct.  The function is implemented by the dog struct.  The function is called when the dog struct is passed to the interface variable.
	fmt.Println(d.Name, "Woof!")
}
func (c Cat) Speak() { // function is created which is used to define the behavior of the cat struct.  The function is implemented by the cat struct.  The function is called when the cat struct is passed to the interface variable.
	fmt.Println(c.Name, "meww")
}
func (d Dog) Move() {
	fmt.Println("Dog is running")
}
func (c Cat) Move() {
	fmt.Println("Cat is jumping")
}

func makeSpeak(s Speaker) {
	s.Speak()
}

func makeMove(s Speaker) {  // for the same task of the diffrent structs the interface is used to define the behavior of the struct.  The interface can be implemented by any struct that has the same method signature.
	s.Move()
}
func main() {

	dog := Dog{Name: "Tommy"}
	cat := Cat{Name: "Kitty"}

	makeSpeak(dog)
	makeSpeak(cat)
	makeMove(dog)   //  calling the same function for the different structs

	makeMove(cat)  ///  calling the same function for the different structs

}
