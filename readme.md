# Go Learning

Chapter-wise markdown files hold the notes; matching folders hold runnable Go code.

```
learning/
├── readme.md                 <- this index
├── chapter-1-introduction.md <- notes
├── chapter-1-introduction/   <- code
│   └── main.go
└── chapter-2-.../            <- next chapters...
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
- Every code folder is a self-contained Go module example you can run with `go run .`