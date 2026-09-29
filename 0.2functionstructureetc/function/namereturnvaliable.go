package main

import "fmt"

// x and y are named return values
func split(sum int) (x, y int) { //   return multiple values from a function (x, y int).
	x = sum * 4 / 9
	y = sum - x
	return // "Naked" return automatically returns x and y
}

func main() {
	fmt.Println(split(22)) // Output: 7 10
}
