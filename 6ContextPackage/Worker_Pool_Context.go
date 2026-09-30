// combining the featuer of the worker pool and the context package.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go worker(ch, i, &wg)
	}
	fmt.Println("Waiting for workers to complete...")
	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
func worker(ch chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Worker", id, "completed the task")
		ch <- fmt.Sprintf("Worker %d finished the task", id)
	case <-ctx.Done():
		fmt.Println("Worker", id, "timed out:", ctx.Err())
	}

}
