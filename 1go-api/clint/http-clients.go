package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	response, err := http.Get("http://localhost:8080/user")  // get request from local host 

	if err != nil {
		fmt.Println("Error:", err)  // handling the error is the server is not running.
		return
	}

	defer response.Body.Close()  // last ma kaam garcha

	body, err := io.ReadAll(response.Body) // io  read the response body and store it in body variable

	if err != nil {
		fmt.Println("Error reading response:", err)   // handel the exception cases if the response body is not readable.
		return
	}

	fmt.Println("Response from Server:")
	fmt.Println(string(body))  
}

