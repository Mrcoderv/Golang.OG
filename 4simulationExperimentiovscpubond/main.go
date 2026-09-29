package main

import (
	"fmt"
	"sync"
	"time"
)

const totalTasks = 100

func main() {

	workerCounts := []int{1, 2, 4, 8, 16, 32}

	cpuTimes := make([]float64, 0, len(workerCounts)) 
	ioTimes := make([]float64, 0, len(workerCounts)) // creating the slices to store execution times

	fmt.Println("Running CPU tests...")

	for _, workers := range workerCounts {

		start := time.Now()

		runCPUWorkerPool(workers, totalTasks)

		elapsed := time.Since(start).Seconds()

		cpuTimes = append(cpuTimes, elapsed)

		fmt.Printf("CPU Workers: %d | Time: %.4f seconds\n", workers, elapsed)
	}

	fmt.Println("\nRunning I/O tests...")

	for _, workers := range workerCounts {

		start := time.Now()

		runIOWorkerPool(workers, totalTasks)

		elapsed := time.Since(start).Seconds()

		ioTimes = append(ioTimes, elapsed)

		fmt.Printf("I/O Workers: %d | Time: %.4f seconds\n", workers, elapsed)
	}

	fmt.Println("\nGenerating graph...")

	err := plotResults(workerCounts, cpuTimes, ioTimes)

	if err != nil {
		panic(err)
	}

	fmt.Println("Graph saved as worker_performance.png")
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

func runIOWorkerPool(numWorkers, totalTasks int) {

	var wg sync.WaitGroup

	taskCh := make(chan int, totalTasks)

	for i := 0; i < numWorkers; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			for id := range taskCh {
				ioTask(id)
			}
		}()
	}

	for i := 1; i <= totalTasks; i++ {
		taskCh <- i
	}

	close(taskCh)

	wg.Wait()
}

