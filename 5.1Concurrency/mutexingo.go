package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {

// 	count := 0
// 	// running  concurrnetly two go routines to increment the count variable

// 	go func() {
// 		for i := 0; i < 100000; i++ {
// 			count++
// 		}
// 		fmt.Println("first 1:", count)
// 	}()

// 	go func() {
// 		for i := 0; i < 100000; i++ {
// 			count++
// 		}
// 		fmt.Println("second 2:", count)
// 	}()

// 	time.Sleep(time.Millisecond * 50)

// 	fmt.Println("final ", count)
// }

// // somertimes this can give the result 11
// //
// // due to overwrite .

// now using the mutex to syncronizse the task

func main() {

	count := 0
	// running  concurrnetly two go routines to increment the count variable
	var mu sync.Mutex
	go func() {
		mu.Lock()
		for i := 0; i < 100000; i++ {
			count++
		}
		fmt.Println("first 1:", count)
		mu.Unlock()
	}()

	go func() {
		mu.Lock() //  blocking the accces to other .

		for i := 0; i < 100000; i++ {
			count++
		}

		fmt.Println("second 2:", count)
		mu.Unlock() // realease the lock.
	}()

	time.Sleep(time.Millisecond * 50)

	fmt.Println("final ", count)
}
