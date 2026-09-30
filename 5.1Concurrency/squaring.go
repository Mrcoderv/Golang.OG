package main

import "fmt"

func square(number int, channel chan int) {
	result := number * number

	channel <- result
}

func main() {
	channel := make(chan int)

	numbers := []int{2, 4, 5, 7, 10}

	for _, number := range numbers {
		go square(number, channel)
	}

	for i := 0; i < len(numbers); i++ {
		result := <-channel
		fmt.Println(result)
	}
}
