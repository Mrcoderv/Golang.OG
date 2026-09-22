package main
import "fmt"
func add(nums ... int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total

}
func greet(name string, greetings ...string) {
	for i, greeting := range greetings {  // _ is used to ignore the index of the slice    like 1 2 3  roll
		fmt.Printf("%d: %s, %s!\n", i, greeting, name)

 
	}
}


func main() {
	result := add(1, 2, 3, 4, 5)
	fmt.Println(result)
// passing a slice to a variadic function
	numbers := []int{10, 20, 30}
	result2 := add(numbers...)
	fmt.Println(result2)

	// passing  other variable on the variadic function
	
    greet("Hello", "Mike", "Liam")   /// here the first name is use ass thegreet string and other as the names using vardic function
	
    greet("Welcome", "Jonathan")


}
