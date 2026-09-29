package main

import (
	"fmt"
)

type Task struct {
	ID int
}
// func main() {
// 	// create a //

//  worker pool with 5 workers
// 	pool := NewWorkerPool(5)

// 	// start the worker pool
// 	pool.Start()

// 	// submit some tasks to the worker pool
// 	for i := 0; i < 10; i++ {
// 		task := &Task{ID: i}
// 		pool.Submit(task)
// 	}
// 	// stop the worker pool
// 	pool.Stop()
// }
// func Task(ID int) *Task {
// 	return &Task{ID: ID}
// }

// // func NewWorkerPool(numWorkers int) *WorkerPool {
// // 	return &WorkerPool{
// // 		numWorkers: numWorkers,
// // 		tasks:      make(chan *Task),
// // 	}
// // }

// // Fan out and fan in patterns

// // use eg file scanning.

// // func worker(n int, results chan int) {
// // 	results <- n * n
// // }

// // func main() {
// // 	results := make(chan int) // shared channel is used to collecte the result.

// // 	// Fan-Out
// // 	go worker(1, results) // dividing the work among multiple workers
// // 	go worker(2, results)
// // 	go worker(3, results)

// // 	// Fan-In
// // 	for i := 0; i < 3; i++ { // collecting the data .
// // 		fmt.Println(<-results)
// // 	}
// // }

// // pipeline pattern
// func main() {
// 	numbers := []int{1, 2, 3, 4, 5} // this is the input data for the pipelune

// 	// create a channel to pass data between stages
// 	stage1 := make(chan int) // first stage channel conecting to second stage channel
// 	stage2 := make(chan int) // second he first stage of the pipeline
// 	go func() {  // level one 
// 		for _, n := range numbers {
// 			stage1 <- n * 2 // double the number
// 		}
// 		close(stage1)
// 	}()
// 	// start the second stage of the pipeline

// 	go func() {
// 		for n := range stage1 {  // here the first stage is connected to the second stage of the pipeline .
// 			stage2 <- n + 1 // add 1 to the number
// 		}
// 		close(stage2)
// 	}()

// 	// collect the results from the second stage of the pipeline
// 	for result := range stage2 {  
// 		fmt.Println(result)
// 	}
// }
// result will be 3, 5, 7, 9, 11
//

// generator pattern
// func generator(nums ...int) <-chan int { // this function will generate a channel of integers
// 	out := make(chan int) // create a channel to send the integers
// 	go func() { // start a goroutine to send the integers to the channel
// 		for _, n := range nums { 
// 			out <- n // send the integer to the channel
// 		}
// 		close(out) // close the channel when done
// 	}()
// 	return out // return the channel
// }

// // generator pattern is used to generate a stream of data that can be consumed by other goroutines. It allows for concurrent processing of data and can be used to implement pipelines and fan-out/fan-in patterns.
// func main() {
// 	nums := []int{1, 2, 3, 4, 5} // this is the input data for the generator

// 	// create a channel to receive the generated integers
// 	ch := generator(nums...) // call the generator function with the input data

// 	// collect the results from the generator
// 	for n := range ch { // receive the integers from the channel
// 		fmt.Println(n) // print the integer
// 	}
// }
// selector pattern


// func main() {
// 	// create two channels to send data to
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	// start a goroutine to send data to the first channel
// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			ch1 <- i
// 		}
// 		close(ch1)
// 	}()

// 	// start a goroutine to send data to the second channel
// 	go func() {
// 		for i := 5; i < 10; i++ {
// 			ch2 <- i
// 		}
// 		close(ch2)
// 	}()

// 	// Wait until ch1 OR ch2 has data available. Whichever is ready, receive from it.

// 	for i := 0; i < 10; i++ {
// 		select {
// 		case n := <-ch1:
// 			fmt.Println("Received from ch1:", n)
// 		case n := <-ch2:
// 			fmt.Println("Received from ch2:", n)
// 		}
// 	}
// }
// bounded patterns
// func main() {
// for i := 0; i < 10; i++ {
	
// 	ch := make(chan int, 5)  // there we descibe the size of buffer .

// 	go func() {
// 		for j := 0; j < 10; j++ {
// 			ch <- j
// 			fmt.Println("Sent:", j)
// 		}
// 		close(ch)
// 	}()


// 	for n := range ch {
// 		fmt.Println("Received:", n)
// 	}
// }
// } 

// producer-consumer pattern
func producer(ch chan<- int) { 
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Produced:", i)
	}
	close(ch)
}

func consumer(ch <-chan int) { 
	for n := range ch {
		fmt.Println("Consumed:", n) 
	}
}

func main() {

	ch := make(chan int, 5)

	go producer(ch) 
	consumer(ch) 
}
	