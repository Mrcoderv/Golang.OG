# Golang.OG

Learning Go, chapter by chapter. Markdown files hold the notes; matching folders hold runnable Go code.

```
.
├── README.md                    <- this index
├── chapter-1-introduction.md    <- notes
├── chapter-1-introduction/      <- code
│   └── main.go
└── chapter-2-.../               <- next chapters...
```

## Chapters

- [Chapter 1: Introduction to Go](./chapter-1-introduction.md)

  Run the example:

  ```
  go run ./chapter-1-introduction
  ```

## Conventions

- Notes live in `chapter-N-short-name.md`
- Code lives in the sibling folder `chapter-N-short-name/`
- Every code folder is a self-contained example you can run with `go run .`

## Official Go Learning Resources

### Go Official Website

The main official website for Go.

### Go Learn — Start Here

Best starting point. It includes documentation, tutorials, Tour of Go, examples, and guided learning.

### A Tour of Go ⭐

My #1 recommendation for you. It's interactive—you write Go code directly in the browser and run it. It covers syntax, data structures, methods, interfaces, generics, and concurrency.

### Official Go Documentation

Use this as your main reference once you know the basics. It covers the language, standard library, modules, testing, concurrency, garbage collection, memory model, etc.

### Official Go Tutorials

Very useful for practical development: modules, JSON, databases, REST APIs, generics, fuzzing, and security tools.

### Effective Go ⭐

After learning the basics, read this to understand how Go code should actually be written—idiomatic Go, error handling, interfaces, concurrency, etc. The official documentation specifically recommends it for new Go programmers.

### Go Standard Library Documentation

Essential when you start building real applications. You can learn packages such as net/http, database/sql, encoding/json, os, sync, context, and more.

## Recommended Learning Order

Since you already have programming experience, I would follow:

```
Go Official Website
       ↓
A Tour of Go
       ↓
Go Tutorials
       ↓
Effective Go
       ↓
Standard Library
       ↓
Concurrency
       ↓
REST API
       ↓
Database
       ↓
Authentication
       ↓
Microservices
       ↓
Docker + Go
       ↓
Production Go
```

Don't start by memorizing Go syntax. Focus especially on structs, interfaces, pointers, error handling, goroutines, channels, context, packages/modules, and net/http. Those are the concepts that will make Go significantly different from Python/Java/JavaScript.

## Useful Go Commands

### Project Setup

```bash
go mod init example.com/myproject   # create a new Go module (creates go.mod)
mkdir myproject && cd myproject     # new project directory (do this first)
```

### Create / Edit Code

```bash
# Create or edit source files in your editor, e.g.:
touch main.go                       # create a new main.go
```

### Run

```bash
go run main.go                      # compile and run a single file
go run .                            # run the package in the current directory
go run ./chapter-1-introduction     # run a package in a subdirectory
```

### Build

```bash
go build .                          # compile into an executable in the current dir
go build -o myapp .                 # build with a custom output name
go build ./...                      # build all packages
GOOS=linux GOARCH=amd64 go build .  # cross-compile for another platform
```

### Format

```bash
gofmt -w main.go                    # format a single file in place
gofmt -w .                          # format everything (check gofmt for dir support)
go fmt ./...                        # format all packages in the project
```

### Test

```bash
go test ./...                       # run all tests
go test -v -run TestName            # run one test verbosely
```

### Dependencies

```bash
go get example.com/some/pkg         # add a dependency
go mod tidy                         # remove unused deps + add missing ones
go mod download                     # download module dependencies
go list -m all                      # list all module dependencies
```

### Check / Fix

```bash
go vet ./...                        # report suspicious constructs
go env                              # show Go environment/config
go version                          # show installed Go version
go doc fmt                          # show package/function docs
go help <command>                   # help for any command
```
