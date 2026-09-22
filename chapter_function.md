# Chapter: Functions in Go — Day 1  
**Learning objectives**  
By the end of this chapter, you should be able to:  
- Declare and call functions.  
- Pass arguments and return one or more values.  
- Use variadic, anonymous, recursive, and higher-order functions.  
- Store functions in variables and pass them as arguments.  
- Define methods and use pointer receivers.  
- Write basic generic functions.  
**1. Basic functions**  
A function is a reusable block of code declared with the func keyword.  
func greet() {  
        fmt.Println("Hello, Go!")  
 }  
   
 func main() {  
        greet()  
 }  
   
**2. Parameters**  
Parameters receive values passed to a function.  
func greet(name string) {  
        fmt.Println("Hello,", name)  
 }  
   
**3. Return values**  
A function can return a value by specifying its return type.  
func add(a, b int) int {  
        return a + b  
 }  
   
**4. Multiple return values**  
Go functions can return multiple values, commonly a result and an error.  
func divide(a, b float64) (float64, error) {  
        if b == 0 {  
               return 0, errors.New("cannot divide by zero")  
        }  
        return a / b, nil  
 }  
   
**5. Variadic functions**  
A variadic function accepts zero or more arguments of the same type. The ...  
   
 syntax must appear before the parameter type.  
func sum(numbers ...int) int {  
        total := 0  
        for _, number := range numbers {  
               total += number  
        }  
        return total  
 }  
   
**6. Anonymous functions**  
An anonymous function has no name and can be called immediately or assigned to  
   
 a variable.  
func() {  
        fmt.Println("Running an anonymous function")  
 }()  
   
**7. Functions as variables**  
Functions are first-class values and can be stored in variables.  
operation := add  
 result := operation(2, 3)  
   
**8. Functions as parameters**  
A function can receive another function as an argument.  
func apply(a, b int, operation func(int, int) int) int {  
        return operation(a, b)  
 }  
   
**9. Functions returning functions**  
A function can return another function.  
func multiplier(factor int) func(int) int {  
        return func(value int) int {  
               return value * factor  
        }  
 }  
   
**10. Recursive functions**  
A recursive function calls itself and must have a stopping condition.  
func factorial(n int) int {  
        if n <= 1 {  
               return 1  
        }  
        return n * factorial(n-1)  
 }  
   
**11. Methods**  
A method is a function associated with a user-defined type.  
type User struct {  
        Name string  
 }  
   
 func (u User) Greet() string {  
        return "Hello, " + u.Name  
 }  
   
**12. Pointer receivers**  
Use a pointer receiver when a method must modify the original value. & gets  
   
 an address and * dereferences a pointer.  
func (u *User) Rename(name string) {  
        u.Name = name  
 }  
   
**13. Generic functions**  
Generic functions use type parameters to work with multiple types.  
func first[T any](items []T) T {  
        return items[0]  
 }  
   
**14. Higher-order functions and **defer  
Higher-order functions accept or return functions. The defer keyword delays a  
   
 function call until the surrounding function returns.  
func main() {  
        defer fmt.Println("Runs last")  
        fmt.Println("Runs first")  
 }  
   
**Day 1 practice tasks**  
1. Write square(n int) int and print the square of 7.  
2.    
3. Write minMax(numbers ...int) (int, int).  
4. Write isEven(n int) bool and pass it to a function that filters numbers.  
5. Create a Rectangle type with an Area() method.  
6. Add a pointer-receiver method that changes a rectangle's width.  
7. Write a generic contains[T comparable] function.  
8. Write a recursive function that calculates the sum from 1 to n.  
9. Use defer to print a message when a function finishes.  
**Checklist**  
- I can declare and call a function.  
- I can use parameters and return values.  
- I can write a variadic function.  
- I understand function values and higher-order functions.  
- I can define methods and pointer receivers.  
- I can write a basic generic function.  
