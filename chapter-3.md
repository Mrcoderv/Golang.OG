# Chapter 3: Variables, Constants, and Data Types

## 2.1 Introduction

A **variable** is a named memory location used to store a value that can change during program execution.

A **constant** is a value whose value cannot be changed after it is declared.

A **data type** defines what kind of value a variable or constant can store, such as an integer, floating-point number, string, or Boolean value.

Go is a **statically typed** language, which means the compiler checks types during compilation.

Example:

```go
package main

import "fmt"

func main() {
    var age int = 21
    var name string = "Ram"
    var passed bool = true

    fmt.Println(age)
    fmt.Println(name)
    fmt.Println(passed)
}
```

Output:

```
21
Ram
true
```

## 2.2 Variables

A variable is a named storage location that holds a value.

**Syntax**

```go
var variableName dataType = value
```

Example:

```go
var age int = 21
```

Here:

- `var` → keyword used to declare a variable
- `age` → variable name
- `int` → data type
- `21` → value

## 2.3 Declaring Variables

There are several ways to declare variables in Go.

**Method 1: Declaration with Type and Value**

```go
var age int = 21
```

Example:

```go
package main

import "fmt"

func main() {
    var age int = 21

    fmt.Println(age)
}
```

Output:

```
21
```

## 2.4 Declaration Without Initial Value

You can declare a variable without assigning a value.

```go
var age int
```

Go automatically assigns its **zero value**.

For an integer, the zero value is:

```
0
```

Example:

```go
package main

import "fmt"

func main() {
    var age int

    fmt.Println(age)
}
```

Output:

```
0
```

## 2.5 Type Inference

Go can automatically determine the type from the assigned value.

```go
var age = 21
```

Go understands that `age` is an integer.

Similarly:

```go
var name = "Ram"
```

Go determines that `name` is a string.

Example:

```go
package main

import "fmt"

func main() {
    var age = 21
    var name = "Ram"
    var height = 5.8

    fmt.Println(age)
    fmt.Println(name)
    fmt.Println(height)
}
```

## 2.6 Short Variable Declaration

Go provides a convenient way to declare variables using `:=`.

**Syntax**

```go
variableName := value
```

Example:

```go
age := 21
name := "Ram"
```

Go automatically determines their types.

```go
package main

import "fmt"

func main() {
    age := 21
    name := "Ram"

    fmt.Println(age)
    fmt.Println(name)
}
```

**Important Rule**

The `:=` operator can be used **inside functions**.

Correct:

```go
func main() {
    age := 21
}
```

Generally incorrect at package level:

```go
age := 21
```

At package level, use:

```go
var age = 21
```

## 2.7 Multiple Variable Declaration

You can declare multiple variables at the same time.

```go
var name, city string
```

You can also assign values:

```go
var name, city = "Ram", "Kathmandu"
```

Example:

```go
package main

import "fmt"

func main() {
    name, age := "Ram", 21

    fmt.Println(name)
    fmt.Println(age)
}
```

Output:

```
Ram
21
```

## 2.8 Reassigning Variables

A variable's value can be changed.

```go
age := 20

age = 21
```

Example:

```go
package main

import "fmt"

func main() {
    age := 20

    fmt.Println(age)

    age = 21

    fmt.Println(age)
}
```

Output:

```
20
21
```

However, you cannot assign a value of an incompatible type.

For example:

```go
age := 20
age = "Ram"
```

This produces a compile-time error because `age` was inferred as an integer.

## 2.9 Constants

A constant is a value that cannot be changed after declaration.

**Syntax**

```go
const constantName = value
```

Example:

```go
const PI = 3.14159
```

Another example:

```go
const country = "Nepal"
```

Example:

```go
package main

import "fmt"

func main() {
    const PI = 3.14159

    fmt.Println(PI)
}
```

Output:

```
3.14159
```

Trying to change it:

```go
const PI = 3.14159

PI = 4
```

will result in a compilation error.

## 2.10 Typed and Untyped Constants

Go supports both typed and untyped constants.

**Typed constant**

```go
const age int = 21
```

The constant has explicitly been given the type `int`.

**Untyped constant**

```go
const age = 21
```

The constant does not have an explicitly declared type.

Go can use an untyped constant in different contexts where the value is representable.

Example:

```go
const x = 10

var a int = x
var b float64 = x

fmt.Println(a)
fmt.Println(b)
```

## 2.11 Constant Expressions

Constants can be created using expressions.

```go
const a = 10
const b = 20
const sum = a + b
```

Example:

```go
package main

import "fmt"

func main() {
    const length = 10
    const width = 5
    const area = length * width

    fmt.Println(area)
}
```

Output:

```
50
```

## 2.12 Multiple Constants

Multiple constants can be declared using a `const` block.

```go
const (
    PI     = 3.14159
    Country = "Nepal"
    Year    = 2026
)
```

Example:

```go
package main

import "fmt"

const (
    PI       = 3.14159
    Language = "Go"
    Version  = 1
)

func main() {
    fmt.Println(PI)
    fmt.Println(Language)
    fmt.Println(Version)
}
```

## 2.13 iota

`iota` is a special identifier used inside constant declarations.

It is useful for creating sequential constants.

Example:

```go
const (
    Monday = iota
    Tuesday
    Wednesday
    Thursday
    Friday
)
```

The values become:

```
Monday     = 0
Tuesday    = 1
Wednesday  = 2
Thursday   = 3
Friday     = 4
```

Example:

```go
package main

import "fmt"

const (
    Monday = iota
    Tuesday
    Wednesday
)

func main() {
    fmt.Println(Monday)
    fmt.Println(Tuesday)
    fmt.Println(Wednesday)
}
```

Output:

```
0
1
2
```

## 2.14 Data Types in Go

Go provides several built-in data types.

The major categories are:

```
Data Types
│
├── Boolean
│
├── Numeric
│   ├── Integer
│   ├── Floating-point
│   └── Complex
│
├── String
│
└── Other Types
    ├── Array
    ├── Slice
    ├── Map
    ├── Struct
    ├── Pointer
    ├── Function
    ├── Interface
    └── Channel
```

In this chapter, we will focus mainly on the basic built-in types.

## 2.15 Boolean Data Type

The Boolean type is:

```
bool
```

It can contain only two values:

```
true
false
```

Example:

```go
var isStudent bool = true
```

Or:

```go
isStudent := true
```

Example:

```go
package main

import "fmt"

func main() {
    var isStudent bool = true

    fmt.Println(isStudent)
}
```

Output:

```
true
```

Boolean values are commonly used with conditions.

```go
age := 20

if age >= 18 {
    fmt.Println("Adult")
}
```

## 2.16 Integer Data Types

Go provides several integer types.

| Type   | Description            |
| ------ | ---------------------- |
| `int`  | Signed integer         |
| `int8` | 8-bit signed integer   |
| `int16`| 16-bit signed integer  |
| `int32`| 32-bit signed integer  |
| `int64`| 64-bit signed integer  |
| `uint` | Unsigned integer       |
| `uint8`| 8-bit unsigned integer |
| `uint16`| 16-bit unsigned integer |
| `uint32`| 32-bit unsigned integer |
| `uint64`| 64-bit unsigned integer |

## 2.17 Signed Integers

Signed integers can contain both positive and negative numbers.

For example:

```go
var age int = 21
var temperature int = -5
```

The `int` type is commonly used for general integer calculations.

## 2.18 int8

`int8` uses 8 bits.

Its range is:

```
-128 to 127
```

Example:

```go
var temperature int8 = -10
```

## 2.19 int16

`int16` uses 16 bits.

Range:

```
-32,768 to 32,767
```

Example:

```go
var number int16 = 30000
```

## 2.20 int32

`int32` uses 32 bits.

Range:

```
-2,147,483,648 to 2,147,483,647
```

Example:

```go
var population int32 = 1000000
```

## 2.21 int64

`int64` uses 64 bits.

Range:

```
-9,223,372,036,854,775,808
to
9,223,372,036,854,775,807
```

Example:

```go
var population int64 = 10000000000
```

## 2.22 Unsigned Integers

Unsigned integers can store only non-negative values.

They use the prefix:

```
u
```

Examples:

```
uint
uint8
uint16
uint32
uint64
```

For example:

```go
var count uint = 100
```

## 2.23 byte

`byte` is an alias for:

```
uint8
```

Therefore:

```go
var b byte = 65
```

is equivalent to:

```go
var b uint8 = 65
```

Example:

```go
package main

import "fmt"

func main() {
    var b byte = 65

    fmt.Println(b)
}
```

Output:

```
65
```

A `byte` is commonly used when working with raw binary data and byte sequences.

## 2.24 rune

`rune` is an alias for:

```
int32
```

It is commonly used to represent a **Unicode code point**.

Example:

```go
var r rune = 'A'
```

You can also use Unicode characters:

```go
var r rune = 'अ'
```

Example:

```go
package main

import "fmt"

func main() {
    var r rune = 'A'

    fmt.Println(r)
}
```

Output:

```
65
```

**Why 65?**

Because 65 is the Unicode code point for `A`.

## 2.25 Floating-Point Data Types

Go provides two floating-point types:

```
float32
float64
```

**float32**

Uses 32 bits.

Example:

```go
var price float32 = 99.50
```

**float64**

Uses 64 bits.

Example:

```go
var price float64 = 99.50
```

`float64` is generally preferred when greater precision is needed.

## 2.26 Complex Numbers

Go supports complex numbers.

Two types are available:

```
complex64
complex128
```

Example:

```go
var c complex64 = 2 + 3i
```

Example:

```go
package main

import "fmt"

func main() {
    var c complex128 = 2 + 3i

    fmt.Println(c)
}
```

Output:

```
(2+3i)
```

## 2.27 String Data Type

A string is a sequence of bytes representing text.

The type is:

```
string
```

Example:

```go
var name string = "Ram"
```

Or:

```go
name := "Ram"
```

Example:

```go
package main

import "fmt"

func main() {
    name := "Ram"

    fmt.Println(name)
}
```

Output:

```
Ram
```

## 2.28 String Concatenation

Strings can be joined using the `+` operator.

```go
firstName := "Ram"
lastName := "Sharma"

fullName := firstName + " " + lastName
```

Example:

```go
package main

import "fmt"

func main() {
    firstName := "Ram"
    lastName := "Sharma"

    fullName := firstName + " " + lastName

    fmt.Println(fullName)
}
```

Output:

```
Ram Sharma
```

## 2.29 String Length

The `len()` function can be used to obtain the number of bytes in a string.

Example:

```go
name := "Ram"

fmt.Println(len(name))
```

Output:

```
3
```

Be careful with Unicode text: `len()` counts **bytes**, not necessarily the number of human-readable characters.

For example, a Unicode character may occupy multiple bytes.

## 2.30 Raw String Literals

Go supports raw string literals using backticks:

```go
`Hello
World`
```

Example:

```go
package main

import "fmt"

func main() {
    message := `Hello
Welcome to Go`

    fmt.Println(message)
}
```

Output:

```
Hello
Welcome to Go
```

Raw strings are useful for:

- Multiline text
- Regular expressions
- File paths
- Text containing quotation marks

## 2.31 Zero Values

One important feature of Go is that variables automatically receive a **zero value** if no initial value is provided.

| Data Type  | Zero Value |
| ---------- | ---------- |
| `int`      | `0`        |
| `float64`  | `0`        |
| `bool`     | `false`    |
| `string`   | `""`       |
| Pointer    | `nil`      |
| Slice      | `nil`      |
| Map        | `nil`      |
| Interface  | `nil`      |
| Function   | `nil`      |
| Channel    | `nil`      |

Example:

```go
package main

import "fmt"

func main() {
    var age int
    var price float64
    var name string
    var active bool

    fmt.Println(age)
    fmt.Println(price)
    fmt.Println(name)
    fmt.Println(active)
}
```

Output:

```
0
0

false
```

## 2.32 Type Conversion

Go does not automatically convert incompatible numeric types.

You must explicitly convert them.

Example:

```go
var x int = 10
var y float64 = float64(x)
```

Here:

```go
float64(x)
```

converts `x` from `int` to `float64`.

Example:

```go
package main

import "fmt"

func main() {
    var x int = 10

    var y float64 = float64(x)

    fmt.Println(x)
    fmt.Println(y)
}
```

Output:

```
10
10
```

## 2.33 Converting Float to Integer

Example:

```go
var x float64 = 10.8
var y int = int(x)
```

The fractional part is discarded.

Output:

```
10
```

Example:

```go
package main

import "fmt"

func main() {
    var x float64 = 10.8

    y := int(x)

    fmt.Println(y)
}
```

Output:

```
10
```

## 2.34 Type Mismatch

Go does not allow incompatible types in an operation.

For example:

```go
var a int = 10
var b float64 = 20.5

// This is invalid:
result := a + b
```

Convert one type first:

```go
result := float64(a) + b
```

Now the calculation is valid.

## 2.35 fmt.Printf() and Format Specifiers

Go provides `fmt.Printf()` for formatted output.

Example:

```go
name := "Ram"
age := 21

fmt.Printf("Name: %s\n", name)
fmt.Printf("Age: %d\n", age)
```

Common format specifiers:

| Specifier | Purpose                        |
| --------- | ------------------------------ |
| `%d`      | Integer                        |
| `%f`      | Floating-point                |
| `%s`      | String                         |
| `%t`      | Boolean                        |
| `%v`      | General/default format         |
| `%T`      | Type of value                  |
| `%c`      | Character/rune                 |

Example:

```go
package main

import "fmt"

func main() {
    age := 21
    price := 99.50
    name := "Ram"

    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Price: %f\n", price)
    fmt.Printf("Name: %s\n", name)
}
```

## 2.36 Finding the Type of a Variable

Use `%T` with `fmt.Printf()`.

```go
package main

import "fmt"

func main() {
    age := 21
    name := "Ram"

    fmt.Printf("%T\n", age)
    fmt.Printf("%T\n", name)
}
```

Output:

```
int
string
```

## 2.37 nil

`nil` represents the absence of a value for certain Go types.

It can be used with types such as:

- Pointers
- Slices
- Maps
- Interfaces
- Functions
- Channels

Example:

```go
var ptr *int

fmt.Println(ptr)
```

Output:

```
<nil>
```

However, `nil` cannot be directly assigned to an ordinary `int`, `float64`, or `bool`.

## 2.38 Variable Scope

The scope of a variable determines where it can be accessed.

**Local Variable**

A variable declared inside a function is local to that function/block.

```go
func main() {
    age := 21

    fmt.Println(age)
}
```

`age` cannot normally be accessed outside its scope.

**Package-Level Variable**

A variable can also be declared outside functions.

```go
package main

import "fmt"

var name = "Ram"

func main() {
    fmt.Println(name)
}
```

Here `name` is declared at package level.

## 2.39 Variable Shadowing

A variable inside a narrower scope can have the same name as a variable from an outer scope.

Example:

```go
package main

import "fmt"

var x = 10

func main() {
    x := 20

    fmt.Println(x)
}
```

Output:

```
20
```

The local `x` shadows the package-level `x`.

## 2.40 Naming Conventions

Go has common naming conventions.

**Variable names**

```
studentName
totalMarks
firstName
```

**Constants**

Constants may commonly use names such as:

```
MaxUsers
DefaultPort
Pi
```

Go generally uses `camelCase` for ordinary multi-word identifiers:

```
studentName
totalMarks
```

Unlike some languages, Go commonly avoids using underscores for ordinary multi-word variable names.

## 2.41 Complete Example

Let's combine variables, constants, and data types.

```go
package main

import "fmt"

const college = "ABC College"

func main() {
    var name string = "Ram"
    age := 21
    var height float64 = 5.8
    var student bool = true

    fmt.Println("College:", college)
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Student:", student)
}
```

Output:

```
College: ABC College
Name: Ram
Age: 21
Height: 5.8
Student: true
```

## 2.42 Variable vs Constant

| Feature       | Variable              | Constant                 |
| ------------- | --------------------- | ------------------------ |
| Keyword       | `var` / `:=`          | `const`                  |
| Value can change? | Yes                | No                       |
| Type          | Has a type            | Can be typed or untyped  |
| Memory behavior | Runtime variable    | Constant value           |
| Example       | `age := 20`           | `const PI = 3.14`        |

Example:

```go
age := 20
age = 21       // Valid
```

But:

```go
const PI = 3.14
PI = 4         // Error
```

## 2.43 Important Data Type Summary

| Type         | Example        |
| ------------ | -------------- |
| `bool`       | `true`         |
| `int`        | `100`          |
| `int8`       | `-10`          |
| `int16`      | `1000`         |
| `int32`      | `100000`       |
| `int64`      | `10000000000`  |
| `uint`       | `100`          |
| `byte`       | `65`           |
| `rune`       | `'A'`          |
| `float32`    | `10.5`         |
| `float64`    | `10.5`         |
| `complex64`  | `2 + 3i`       |
| `complex128` | `2 + 3i`       |
| `string`     | `"Hello"`      |

## 2.44 Important Exam Points

Remember these points:

- `var` is used to declare variables.
- `:=` is used for short variable declaration.
- `:=` is generally used inside functions.
- `const` is used to declare constants.
- A variable's value can be changed.
- A constant's value cannot be changed.
- Go is statically typed.
- Go supports type inference.
- Go does not automatically perform many incompatible type conversions.
- Explicit conversion uses syntax such as `float64(x)`.
- `bool` has values `true` and `false`.
- `byte` is an alias for `uint8`.
- `rune` is an alias for `int32`.
- `float32` and `float64` are floating-point types.
- `complex64` and `complex128` are complex-number types.
- The zero value of `int` is `0`.
- The zero value of `bool` is `false`.
- The zero value of `string` is an empty string `""`.
- `iota` is used to generate sequential constant values.
- Go is case-sensitive.