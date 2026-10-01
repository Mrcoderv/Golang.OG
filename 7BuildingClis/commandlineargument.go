// package main

// import (
// 	"flag"
// 	"fmt"
// )

// func main() {
// 	name := flag.String("name", "Guest", "Your name")
// 	age := flag.Int("age", 0, "Your age")

// 	flag.Parse()

// 	fmt.Println("Name:", *name)
// 	fmt.Println("Age:", *age)
// }
// //flag.Parse() is used to break the  the command-line flags provided by the user. It processes the flags defined using flag.String, flag.Int, and other flag functions, and assigns the values to the corresponding variables. After calling flag.Parse(), you can access the values of the flags through the pointers returned by the flag functions.
// now this fragment of code is used to create a command-line application that accepts two flags: name and age. The flag.String function defines a string flag named "name" with a default value of "Guest" and a description. The flag.Int function defines an integer flag named "age" with a default value of 0 and a description. After parsing the flags, the program prints the values of the name and age flags to the console.