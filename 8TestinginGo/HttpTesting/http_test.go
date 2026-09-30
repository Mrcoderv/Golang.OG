package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {  // TestHelloHandler is a test function that tests the helloHandler function.
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)

	rec := httptest.NewRecorder()

	helloHandler(rec, req)

	if rec.Code != http.StatusOK {  // comparing the status code with the expected status code.
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code) 
	}

	if rec.Body.String() != "Hello Raghav" {
		t.Errorf("unexpected response: %s", rec.Body.String())
	}
}
