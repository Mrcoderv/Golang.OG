package main

import "fmt"

func call() {
	fmt.Println("Welcome! to raghav's world of golang")
	func() { // anonymous function  // Anonymous function   // directly run when the outside function is called
		// through the braces ()
		// useusss data base connection and other things

		fmt.Println("This is an anonymous function")
	}() // calling the anonymous function
	}
func GFG() func(i, j string) string {
    myf := func(i, j string) string {
        return i + j + "GeeksforGeeks"
    }
    return myf
}

func main() {
	call() // calling the function
	func(ele string) {
		fmt.Println(ele)
	}("raghav ") // passing the value to the anonymous function

    value := GFG()
    fmt.Println(value("Welcome ", "to "))  /// here the value contain the function then  then we pass the value to function /



	/*
	GFG()
 │
 │ returns
 ▼
myf
 │
 │ stored in
 ▼
value
 │
 │ called with
 ▼
value("Welcome ", "to ")
 │
 ▼
"Welcome to RAGHAV"
*/

}
