package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", userHandler)
	http.HandleFunc("/search", searchHandler)

	log.Println("server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// /search demonstrates reading a query parameter.
// Example: GET /search?q=golang
func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "missing q query parameter", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "search query: %s\n", query)
}

// /users demonstrates GET with query parameters and POST.
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Example: GET /users?name=ram&page=2
		name := r.URL.Query().Get("name")
		page := r.URL.Query().Get("page")

		fmt.Fprintf(w, "GET users: name=%q, page=%q\n", name, page)

	case http.MethodPost:
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "POST created user: %+v\n", user)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// /users/{id} demonstrates PUT, PATCH, DELETE, HEAD, and OPTIONS.

func userHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]
	if id == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		fmt.Fprintf(w, "PUT replaced user %s\n", id)

	case http.MethodPatch:
		fmt.Fprintf(w, "PATCH updated user %s\n", id)

	case http.MethodDelete:
		fmt.Fprintf(w, "DELETE removed user %s\n", id)

	case http.MethodHead:
		w.Header().Set("X-User-ID", id)

	case http.MethodOptions:
		w.Header().Set("Allow", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
