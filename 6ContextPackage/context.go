package main

import (
	"context"
	"fmt"
	"time"
)

// 1. // time out cotext example
// func main() {

// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)  // stop on 2 second
// 	defer cancel()

//		select {   // operation is  complete or not .
//		case <-time.After(3 * time.Second):
//			fmt.Println("Operation completed")
//		case <-ctx.Done():
//			fmt.Println("Operation timed out:", ctx.Err())
//		}
//	}
//
// // 2. // cancel context example
// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	go worker(ctx)

// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("Worker stopped:", ctx.Err())

// 	case <-time.After(5 * time.Second):
// 		fmt.Println("Worker completed")
// 		cancel()
// 	}
// }
// func worker(ctx context.Context) {
// 	fmt.Println("Worker started")

// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("Worker received cancellation")
// 		return

// 	case <-time.After(3 * time.Second):
// 		fmt.Println("Worker finished its task")
// 	}
// }

// // context.TODO()
// func main() {
// 	ctx := context.TODO() // Create a context.TODO() context

// 	// Simulate some work
// 	select {
// 	case <-time.After(2 * time.Second):
// 		fmt.Println("Work completed")
// 	case <-ctx.Done():
// 		fmt.Println("Context canceled:", ctx.Err())

// 	}}
// // }
//)2/ // context.withdeadline()
// func main() {
// 	deadline := time.Now().Add(2 * time.Second) // Set a deadline 2 seconds from now
// 	ctx, cancel := context.WithDeadline(context.Background(), deadline)
// 	defer cancel()

// 	select {
// 	case <-time.After(3 * time.Second):
// 		fmt.Println("Work completed")
// 	case <-ctx.Done():
// 		fmt.Println("Context deadline exceeded:", ctx.Err())
// 	}
// }

// context.WithValue() example   this is used to return the value  of code .
// //
// func main() {
// 	ctx := context.WithValue(
// 		context.Background(),
// 		"userID",
// 		123,
// 	)

// 	id := ctx.Value("userID")

// 	fmt.Println(id)

// }

// context.withcancel() example
// func main() {
// 	// Create a cancellable context
// 	ctx, cancel := context.WithCancel(context.Background()) /// here we manually cancell the task

// 	// Start a goroutine that does some work
// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Println("Goroutine received cancellation signal")
// 				return
// 			default:
// 				// Simulate work
// 				fmt.Println("Goroutine is working...")
// 			}
// 		}
// 	}()

// 	// Simulate some work in the main goroutine
// 	//checkin the canceel context

// 	time.Sleep(4 * time.Second)
// 	cancel() //checking the cancel context. after sleep

// 	// Cancel the context

// }

nw
