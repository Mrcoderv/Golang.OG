// This is a simple example of using a function with variadic parameters.
package main

import "fmt"

func minMax(numbers ...int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	return min, max
}

func main() {
	min, max := minMax(10, 5, 20, 3, 15)

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
