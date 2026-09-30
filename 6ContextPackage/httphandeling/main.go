package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Cleanup crew

	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.example.com", nil)
	if err != nil {
		fmt.Println("error aagayo ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("BOO:", err) // Times out? You’ll see context.DeadlineExceeded
		return
	}
	defer resp.Body.Close()

	fmt.Println("BOOM", resp.Status) // we get response.
}
