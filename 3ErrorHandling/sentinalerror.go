// pre define d error that are compared to the latest error condition

// like think like a dictionary of the error that are predefined and can be used to compare with the latest error condition
package main

import (
	"errors"
	"fmt"
)

// predefined errors
var ErrNotFound = errors.New("item not found") // heree we costomize the error message and use it as a predefined error that can be used to compare with the latest error condition

func findItem() error {
	return ErrNotFound
}

func main() {
	err := findItem()

	if err == ErrNotFound {
		fmt.Println("Item was not found")
	}
}
