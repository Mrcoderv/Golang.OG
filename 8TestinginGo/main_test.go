package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)

	if result != 5 {
		t.Errorf("expected 5, got %d", result) // t is used to report errors in the test  ..
		//  like the print function of the error.
	}
}


