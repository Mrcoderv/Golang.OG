# Chapter 2: Go Project Format, Packages, and Imports

## 2.1 Go Project Format (Go Modules)

A modern Go project is a **module**. A module is a collection of related Go packages, managed with a `go.mod` file at the root.

```
chapter-2/
│
├── go.mod          <- module definition
├── main.go         <- main package (program entry point)
└── util/
    └── util.go     <- local package
```

### go.mod

`go.mod` is created with:

```bash
go mod init golang.og/chapter-2
```

It defines the **module path** — the "address" used in imports:

```go
module golang.og/chapter-2

go 1.27.0
```

> Note: pick a valid module path. Names like `main.go` or `util.go` are NOT valid — they end in `.go`, which Go treats as a file, not a module path.

## 2.2 Package and Directory Connection

In Go, **one directory = one package**.

- All Go files in the same folder belong to the same package.
- The package name may differ from the folder name, but it is best practice to make them match.

```
chapter-2/
├── main.go        -> package main
└── util/
    └── util.go    -> package util
```

Every Go file begins with its package declaration:

**main.go**

```go
package main
```

**util/util.go**

```go
package util
```

### Two special packages

| Package     | Meaning                                               |
| ----------- | ----------------------------------------------------- |
| `main`      | Builds into an executable program (needs `func main()`) |
| others      | Build into reusable code that is imported by other packages |

## 2.3 Loading a Local Package (Importing)

To use code from another directory inside the same module, **import it using the full module path**.

`main.go`:

```go
package main

import (
	"golang.og/chapter-2/util"
)

func main() {
	util.PrintMessage("Hello, Raghav")
}
```

The import path is built as:

```
module-path + folder-path-from-root

golang.og/chapter-2            (module path from go.mod)
             + util            (folder name)
             = golang.og/chapter-2/util
```

`util/util.go` exports the function (note the capital `P`):

```go
package util

import "fmt"

func PrintMessage(message string) {
	fmt.Println(message)
}
```

Only **exported** names (starting with an uppercase letter) can be used from other packages. Uppercase = exported.

### Example: adding a helper to the util package

Add `util/math.go`:

```go
package util

func Add(a int, b int) int {
	return a + b
}
```

Use it from `main.go`:

```go
import (
	"golang.og/chapter-2/util"
)

func main() {
	util.PrintMessage("Hello, Raghav")
	util.PrintMessage(fmt.Sprint(util.Add(2, 3)))
}
```

## 2.4 Automatic Importing Feature (gopls)

When you type `util.PrintMessage(...)` in VS Code, the import line appears automatically:

```go
import (
	"golang.og/chapter-2/util"
)
```

### How it works

1. The **Go extension** (`golang.go`) runs the language server **gopls** in the background.
2. As you type, gopls watches your `go.mod` module path: `golang.og/chapter-2`.
3. When you use a symbol it can't find locally, gopls searches the project for the package (here, the `util/` folder) and builds the import path:

   `module path + folder path  =  golang.og/chapter-2/util`

4. gopls inserts the import for you automatically.

### Requirements for auto-import to work

- The **Go extension** `golang.go` must be installed.
- `gopls` must be installed: `go install golang.org/x/tools/gopls@latest`
- The project must have a valid `go.mod`.
- The code folder must belong to the module (no nested second `go.mod`).

## 2.5 Common Mistakes

| Mistake                    | Problem                                        | Fix                              |
| -------------------------- | ---------------------------------------------- | -------------------------------- |
| `module main.go`           | "main.go" is not a valid module path           | use `golang.og/chapter-2`        |
| Nested `util/go.mod`       | creates a second, separate module              | delete it — one module per project |
| `import "mutil"`           | no package named `mutil` at that path          | import the real path + folder    |
| `package mutil` in util.go | package name doesn't match its use             | use `package util`               |

## 2.6 Summary

- One **directory = one package**.
- `go.mod`'s module path is the prefix of every import.
- Local packages are loaded with `module-path/folder-path`.
- gopls + the Go extension add imports automatically while you type.
- A `main` package with `func main()` produces an executable; everything else is reusable library code.