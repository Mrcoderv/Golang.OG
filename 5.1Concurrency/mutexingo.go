package main

import (
	"fmt"
	"time"
)

func main() {

	count := 10
	// running  concurrnetly two go routines to increment the count variable

	go func() {
		count = count * count
		fmt.Println("first 1:", count)
	}()

	go func() {
		count = count * count
		fmt.Println("second 2:", count)
	}()

	time.Sleep(time.Millisecond * 50)

	fmt.Println("final ", count)
}

// // somertimes this can give the result 11
// //
// // due to overwrite .

// now using the mutex to syncronizse the task

// func main() {

// 	count := 10
// 	// running  concurrnetly two go routines to increment the count variable
// 	var mu sync.Mutex
// 	go func() {
// 		mu.Lock()
// 		count = count * count
// 		fmt.Println("first 1:", count)
// 		mu.Unlock()
// 	}()

// 	go func() {
// 		mu.Lock() //  blocking the accces to other .
// 		count = count * count
// 		fmt.Println("second 2:", count)
// 		mu.Unlock() // realease the lock.
// 	}()

// 	time.Sleep(time.Millisecond * 50)

// 	fmt.Println("final ", count)
// }
