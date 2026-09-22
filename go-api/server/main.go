package main

import (
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

func hello(w http.ResponseWriter, r *http.Request) { // main page localhost:8080/
	fmt.Fprintln(w, "Hello from Go API!")
}

func user(w http.ResponseWriter, r *http.Request) { // r is htttp  request.
	u := User{ // creating user object.
		Name: "Raghav",
		Age:  20,
		Address: Address{
			City:    "Gulmi Tamghas",
			Country: "Nepal",
		},
	}

	fmt.Fprintf(w, "Name: %s\nAge: %d\nCity: %s\nCountry: %s", u.Name, u.Age, u.Address.City, u.Address.Country) // w is https response writer, r is http request
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/user", user)

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil) // listen on port 8080  // sumtimes it may stop working without listeninng.

}
