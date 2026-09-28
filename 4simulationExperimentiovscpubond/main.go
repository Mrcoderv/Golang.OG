package main

import (
	"fmt"
	"sync"
	"time"
)



const totalTasks = 100

func main() {
	var totalworkers int
	fmt.Print("enter number of workers cpu ")
	fmt.Scan(&totalworkers)

	start := time.Now()

	runCPUWorkerPool(totalworkers, totalTasks)

	fmt.Println("cpu time", time.Since(start))
}

func runCPUWorkerPool(numWorkers, totalTasks int) {

	var wg sync.WaitGroup

	taskCh := make(chan int, totalTasks)

	for i := 0; i < numWorkers; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			for id := range taskCh {
				cpuTask(id)
			}
		}()
	}

	for i := 1; i <= totalTasks; i++ {
		taskCh <- i
	}
	close(taskCh)

	wg.Wait()
}
