package main

import "fmt"

func main() {

	count := 0

	increment := func() {   // this is a closure function which can access the variable of the outer function 
		count++  
		fmt.Println(count)
	}

	increment()
	increment()
	increment()
}
