package main

import (
	"fmt"
	"sync"
	"time"
)

// CONTEXT CANCELLATION

// func worker(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("worker stopped due to context cancellation")
// 			return

// 		default:
// 			fmt.Println("worker is running")
// 			time.Sleep(900 * time.Millisecond)

// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go worker(ctx)

// 	time.Sleep(2 * time.Second) // Wait for the worker to run for a while
// 	cancel()

// 	time.Sleep(1 * time.Second) // Wait for the worker to finish
// }

// CONTEXT CANCELLATION
// func worker(jobs <-chan int) {
// 	for job := range jobs {
// 		fmt.Println("processing job:", job)
// 	}
// 	fmt.Println("worker stopped")

// }
// func main() {
// 	jobs := make(chan int)

// 	go worker(jobs)

// 	jobs <- 1
// 	jobs <- 2
// 	jobs <- 3

// 	close(jobs)
// 	time.Sleep(1 * time.Second) // Wait for the worker to finish
// }

// WaitGroup / errgroup

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done() // Decrement the counter when the goroutine completes

// 	fmt.Printf("Worker %d is starting\n", id)
// 	time.Sleep(1 * time.Second) // Simulate work
// 	fmt.Printf("Worker %d is done\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	numWorkers := 8
// 	wg.Add(numWorkers) // Set the number of goroutines to wait for

// 	for i := 1; i <= numWorkers; i++ {
// 		go worker(i, &wg) // Start a worker goroutine
// 	}

// 	wg.Wait() // Wait for all workers to finish
// 	fmt.Println("All workers are done")
// }

// using the Bounded concurrency/limited worker

// func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done() // Decrement the counter when the goroutine completes

// 	for job := range jobs {
// 		fmt.Printf("Worker %d is processing job: %d\n", id, job)
// 		time.Sleep(1 * time.Second) // Simulate work
// 	}
// 	fmt.Printf("Worker %d is done\n", id)
// }

// func main() {
// 	numWorkers := 3
// 	numJobs := 10

// 	jobs := make(chan int, numJobs)
// 	var wg sync.WaitGroup

// 	// Start worker goroutines
// 	for i := 1; i <= numWorkers; i++ {
// 		wg.Add(1)
// 		go worker(i, jobs, &wg)
// 	}

// 	// Send jobs to the workers
// 	for j := 1; j <= numJobs; j++ {
// 		jobs <- j
// 	}
// 	close(jobs) // Close the jobs channel to signal no more jobs

// 	wg.Wait() // Wait for all workers to finish
// 	fmt.Println("All workers are done")
// }


// Explicit exit 
package main

import (
	"fmt"
	"time"
)

func worker() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Worker processing:", i)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Worker exiting")
}

func main() {
	go worker()

	time.Sleep(3 * time.Second)

	fmt.Println("Main finished")
}