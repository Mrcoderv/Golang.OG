//Create a parent context and a child context. Cancel the parent and demonstrate that the child is also cancelled.4

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx, cancelParent := context.WithCancel(context.Background())
	childCtx, cancelChild := context.WithCancel(parentCtx)
	go func() {
		select {
		case <-parentCtx.Done():
			fmt.Println("Parent context cancelled")
			cancelChild() // Cancel the child context when the parent is cancelled
		}
	}()

	go func() {
		select {
		case <-childCtx.Done():
			fmt.Println("Child context cancelled")
		}
	}()

	time.Sleep(1 * time.Second)
	cancelParent() // Cancel the parent context

	time.Sleep(1 * time.Second) // Wait for goroutines to finish
}
