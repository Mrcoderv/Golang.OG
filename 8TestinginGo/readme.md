Testing is the process of checking whether the system works as expected in Go.
The main purpose is to check whether the current behavior of a function or method is working correctly or not.
// 
package main

func Add(a, b i nt) int {
	return a + b
}

// test //
package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)

	if result != 5 {
		t.Errorf("expected 5, got %d", result)  // t is used to report errors in the test  ..
		//  like the print function of the error.
	}
}


