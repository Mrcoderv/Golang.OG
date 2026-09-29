# Go HTTP API Notes

This project contains a small HTTP server and a client that communicate over
`http://localhost:8080`.

## `server/main.go`

### Data structures

- `Address` stores a user's city and country.
- `User` stores the user's name, age, and nested `Address` value.

### Routes

The server registers two handlers:

- `GET /` calls `hello` and returns `Hello from Go API!`.
- `GET /user` calls `user` and returns a formatted user record:

	```text
	Name: Raghav
	Age: 20
	City: Gulmi Tamghas
	Country: Nepal
	```

The handler receives an `http.ResponseWriter` (`w`) to write the response and
an `*http.Request` (`r`) containing the incoming request.

### Starting the server

`main` registers the routes and starts an HTTP server on port `8080`:

```go
http.ListenAndServe(":8080", nil)
```

The server must be running before the client is started.

## `clint/http-clients.go`

The client sends a `GET` request to the server's `/user` endpoint:

```go
response, err := http.Get("http://localhost:8080/user")
```

Its flow is:

1. Send the request with `http.Get`.
2. Print an error and stop if the request fails.
3. Defer `response.Body.Close()` so the response body is released.
4. Read the body with `io.ReadAll`.
5. Print the response as a string.

The client does not decode JSON because the server currently returns plain text
using `fmt.Fprintf`.

## Running the example

Open one terminal and start the server:

```bash
cd server
go run .
```

In a second terminal, run the client:

```bash
cd clint
go run .
```

The client should print the user information returned by `GET /user`.
