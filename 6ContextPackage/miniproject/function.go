// Create a function:

// func doWork(ctx context.Context)

// Inside it, simulate work for 5 seconds. If the context is cancelled before the work finishes, stop the function.

package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(2* time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Work cancelled")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go doWork(ctx)

	time.Sleep(3 * time.Second) // Wait for the work to finish or be cancelled
	
}
