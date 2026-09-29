//
# Modules & Dependencies in Go

In Go, **modules** are used to manage your project and its dependencies (external packages/libraries).

Think of it like:

```text
Go Project
   │
   ├── Your code
   │
   ├── Go Module
   │      └── go.mod
   │
   └── Dependencies
          ├── Package A
          ├── Package B
          └── Package C
```

---

## 1. What is a Go Module?

A **module** is a collection of Go packages that are versioned and managed together.

For example, suppose you create a project:

```text
myapp/
├── main.go
└── go.mod
```

The `go.mod` file tells Go:

* What your module is called
* Which Go version you're using
* Which external dependencies your project needs
* Which versions of those dependencies are required

---

# 2. Creating a Module

Create a project:

```bash
mkdir myapp
cd myapp
```

Initialize a Go module:

```bash
go mod init myapp
```

You will get:

```text
myapp/
├── go.mod
```

The `go.mod` file might contain:

```go
module myapp

go 1.25
```

The exact Go version depends on your installed Go version.

---

# 3. `module` directive

Inside `go.mod`:

```go
module myapp
```

This is the **module path**.

For a GitHub project, you would usually use the repository path:

```go
module github.com/username/myapp
```

For example:

```go
module github.com/Mrcoderv/myapp
```

Then your own packages can be imported using:

```go
import "github.com/Mrcoderv/myapp/something"
```

---

# 4. What is a Dependency?

A **dependency** is an external package that your project uses.

For example, suppose you want to use a UUID library:

```go
github.com/google/uuid
```

Your project depends on that package.

Conceptually:

```text
Your Application
       ↓
   UUID package
       ↓
   Other packages
```

---

# 5. Adding a Dependency

You can use:

```bash
go get github.com/google/uuid
```

Go downloads the dependency and updates your `go.mod`.

It may add something like:

```go
require github.com/google/uuid v1.6.0
```

The exact version may differ.

---

# 6. Using the Dependency

Example:

```go
package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()

	fmt.Println(id)
}
```

Run:

```bash
go run .
```

You might get:

```text
550e8400-e29b-41d4-a716-446655440000
```

---

# 7. `go.mod`

A typical `go.mod` might look like:

```go
module github.com/Mrcoderv/myapp

go 1.25

require github.com/google/uuid v1.6.0
```

### Meaning

```go
module github.com/Mrcoderv/myapp
```

Your module name.

```go
go 1.25
```

Go language/toolchain version declaration.

```go
require github.com/google/uuid v1.6.0
```

Your project requires the UUID package at that version.

---

# 8. `go.sum`

When you download dependencies, Go may also create:

```text
go.sum
```

So your project becomes:

```text
myapp/
├── go.mod
├── go.sum
└── main.go
```

### `go.mod`

Describes your module and dependencies.

### `go.sum`

Contains cryptographic checksums that help Go verify downloaded module content.

Think:

```text
go.mod
   ↓
"What dependencies and versions do I need?"

go.sum
   ↓
"Are the downloaded dependencies exactly what they should be?"
```

---

# 9. `go mod tidy`

One of the most useful commands:

```bash
go mod tidy
```

It:

* Adds missing dependencies
* Removes unused dependencies
* Updates `go.mod`
* Updates `go.sum`

For example, if you remove a package from your code and run:

```bash
go mod tidy
```

Go can remove the dependency if it is no longer required.

---

# 10. Download dependencies

You can use:

```bash
go mod download
```

This downloads the modules required by your project.

Usually, you don't need to manually run it because commands like:

```bash
go run .
```

or:

```bash
go build
```

can automatically download required dependencies.

---

# 11. Verify dependencies

You can use:

```bash
go mod verify
```

This verifies that dependencies haven't been modified after being downloaded.

---

# 12. List dependencies

You can see dependencies using:

```bash
go list -m all
```

Example:

```text
github.com/Mrcoderv/myapp
github.com/google/uuid v1.6.0
```

---

# 13. Dependency versions

Go modules use versions.

For example:

```go
require github.com/google/uuid v1.6.0
```

means your project is using:

```text
uuid
version 1.6.0
```

You can upgrade a dependency with:

```bash
go get -u github.com/google/uuid
```

Then:

```bash
go mod tidy
```

---

# 14. Module vs Package

This is an important exam concept.

### Package

A **package** is a collection of related Go source files.

Example:

```text
utils/
├── math.go
└── string.go
```

These files might belong to:

```go
package utils
```

### Module

A **module** is a collection of packages managed together.

```text
Module
│
├── package main
├── package utils
├── package database
└── package models
```

So:

```text
Module
  ↓
contains packages
  ↓
packages contain Go files
```

---

# 15. Example Project Structure

A larger Go application might look like:

```text
myapp/
│
├── go.mod
├── go.sum
│
├── main.go
│
├── models/
│   └── user.go
│
├── database/
│   └── database.go
│
└── utils/
    └── helper.go
```

`go.mod`:

```go
module github.com/Mrcoderv/myapp

go 1.25

require (
	github.com/google/uuid v1.6.0
)
```

---

# 16. Important Go Module Commands

| Command           | Purpose                            |
| ----------------- | ---------------------------------- |
| `go mod init`     | Create a new module                |
| `go get`          | Add/update a dependency            |
| `go mod tidy`     | Clean and synchronize dependencies |
| `go mod download` | Download dependencies              |
| `go mod verify`   | Verify dependencies                |
| `go list -m all`  | List modules                       |
| `go mod graph`    | Show dependency graph              |

---

# 17. Dependency Graph

Suppose:

```text
Your App
   ↓
Library A
   ↓
Library B
```

Then:

```text
Your App
   │
   └── Library A
          │
          └── Library B
```

These are called **transitive dependencies**.

You directly use Library A, but Library A itself depends on Library B.

Go's module system manages these dependencies and their versions.

---

# 18. Simple Real Example

Create:

```bash
mkdir hello
cd hello
go mod init hello
```

Create `main.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go")
}
```

Run:

```bash
go run .
```

You now have:

```text
hello/
├── go.mod
└── main.go
```

Then add an external dependency:

```bash
go get github.com/google/uuid
```

Now your module manages that dependency through `go.mod`.

---

## ⭐ Exam Definition

> **A Go module is a collection of related Go packages that are versioned and managed together. Dependencies are external modules or packages required by a Go project. Go uses `go.mod` and `go.sum` to manage dependencies and their versions.**

### Remember the core relationship:

```text
Go Module
    ↓
   go.mod
    ↓
Dependencies
    ↓
External Packages
```

And the most important commands:

```bash
go mod init
go get
go mod tidy
go mod download
go mod verify
```

**For modern Go development, `go.mod` is the central file for module and dependency management.**
