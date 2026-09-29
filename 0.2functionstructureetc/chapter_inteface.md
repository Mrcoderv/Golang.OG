
# Chapter: Interfaces in Go

Interfaces are one of Go's most useful tools for writing flexible and reusable
code. An interface describes **what a type can do**, rather than **what a type
is**.

## 1. Declaring an interface

An interface contains method signatures:

```go
type Speaker interface {
	Speak() string
}
```

Any type that has a `Speak() string` method automatically satisfies
`Speaker`. Go uses **implicit interface implementation**; there is no
`implements` keyword.

## 2. Implementing an interface

```go
package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, I am " + p.Name
}

func announce(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	person := Person{Name: "Maya"}
	announce(person) // Person satisfies Speaker.
}
```

The `announce` function accepts any value that satisfies `Speaker`. It does
not need to know about the concrete type `Person`.

## 3. Multiple implementations

Different types can satisfy the same interface:

```go
type Robot struct{}

func (Robot) Speak() string {
	return "Beep boop"
}

func main() {
	announce(Person{Name: "Maya"})
	announce(Robot{})
}
```

This makes interfaces useful for replacing implementations without changing
the code that uses them.

## 4. The empty interface and `any`

The empty interface has no methods, so every value satisfies it:

```go
var value any = 42
```

`any` is an alias for `interface{}`. It should be used only when a function
really must accept values of any type; a more specific interface usually gives
better compile-time safety.

## 5. Type assertions

A type assertion retrieves the concrete value stored in an interface:

```go
value := any("hello")

text, ok := value.(string)
if ok {
	fmt.Println(text)
}
```

The comma-ok form is safe. A direct assertion, such as `value.(string)`,
panics if the value is not a string.

## 6. Type switches

Use a type switch when different concrete types need different behavior:

```go
func describe(value any) {
	switch v := value.(type) {
	case string:
		fmt.Println("string:", v)
	case int:
		fmt.Println("integer:", v)
	default:
		fmt.Println("unknown type")
	}
}
```

## 7. Nil interfaces

An interface is `nil` only when both its dynamic type and dynamic value are
`nil`:

```go
var speaker Speaker
fmt.Println(speaker == nil) // true

var person *Person
speaker = person
fmt.Println(speaker == nil) // false: it contains a *Person type
```

This distinction is important when returning pointers through interfaces.

## 8. Pointer receivers

A method with a pointer receiver belongs to the pointer type, not the value
type:

```go
func (p *Person) Speak() string {
	return "Hello, I am " + p.Name
}
```

In this case, `*Person` satisfies `Speaker`, while `Person` does not. Use a
pointer receiver when the method must modify the value or when copying it is
undesirable.

## 9. Practical guidelines

- Define small interfaces, often with one or two methods.
- Define interfaces where they are consumed, not necessarily where types are
  implemented.
- Accept interfaces when a function needs behavior; return concrete types when
  callers need the full implementation.
- Prefer compile-time checks and avoid unnecessary type assertions.

Interfaces let code depend on behavior instead of concrete details, making
programs easier to test, extend, and maintain.
