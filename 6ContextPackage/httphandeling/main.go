package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	url := "https://api.example.com"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("error comes ", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("nope :", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("yup ", resp.Status)
}
