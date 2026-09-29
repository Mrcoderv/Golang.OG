// is used to syncronize the access of the 	shared resources of the go routine s
package main

import (
	"fmt"
	"sync"
)

var (
	x  int
	wg sync.WaitGroup
	mu sync.Mutex
)

func main() {
	wg.Add(2)
	go increment("first")
	go increment("second")
	wg.Wait()
	fmt.Println("Final value of x:", x)
}

func increment(s string) {
	for i := 0; i < 20; i++ {
		mu.Lock()
		x++
		fmt.Println(s, i, "Value of x:", x)
		mu.Unlock()
	}
	wg.Done()
}
