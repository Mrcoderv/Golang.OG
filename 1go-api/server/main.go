package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Address struct {
	City    string
	Country string
}

type User struct {
	Name    string
	Age     int
	Address Address // nested struct.
}

// func (w http.ResponseWriter, r *http.Request) { // main page localhost:8080/
// 	fmt.Fprintln(w, "Hello from Go API!")
// }

func userHandler(w http.ResponseWriter, r *http.Request) { // r is htttp  request.
	u := User{ // creating user object.
		Name: "Raghav",
		Age:  20,
		Address: Address{
			City:    "Gulmi Tamghas",
			Country: "Nepal",
		},
	}
	// data, err := json.Marshal(u)
	// if err != nil {
	// 	http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
	// 	return
	// }

	data, err := xml.Marshal(u)  // converting to xml format.

	if err != nil {
		http.Error(w, "Error marshaling XML", http.StatusInternalServerError)
		return
	}

	// Set the content type of the response to XML
	w.Header().Set("Content-Type", "application/xml")

	w.Write(data)

	//w.Header().Set("Content-Type", "application/json") // setting the content type of the response to JSON.

	// w.Header().Set("Content-Type", "application/json") // setting the content type of the response to JSON.
	// w.Write(data)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // main page localhost:8080/
		fmt.Fprintln(w, "Hello from Go API!")
	})
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil) // listen on port 8080  // sumtimes it may stop working without listeninng.
}
