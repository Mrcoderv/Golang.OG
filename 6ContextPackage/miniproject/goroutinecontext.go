//Create a goroutine that takes 5 seconds to finish,
//but use a context with a 2-second timeout. Print "Cancelled" if the context expires.

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

// 	defer cancel()

// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		fmt.Println("Finished")
// 	}()

// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("Cancelled")
// 	}
// }
