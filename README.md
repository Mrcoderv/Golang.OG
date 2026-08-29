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