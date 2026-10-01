// first order function are treated like thhe value in go . so that teheuy are assign with the variabe

func main() {
	add := func(a int, b int) int {  
		return a + b
	}

	
	result := add(5, 3)
	fmt.Println("Result:", result) 
}


// higher order function can take a function as an argument or return a function.

// so hof is mainly used to use the same block of code for the different operationn
.. like the ently lobby will be same so that other thing can be achive .
think a like the kitchen where different type of order are prepared
so for the seperated food we donot want differnt kitchen

so the higher order function .  has the two approach 
one === consuminbg the function throgh the parameter and 

two == returning the function as the return value .


now by combining the generic T  and interface and the higher-order functions  we cna create the higher order function.


//// closure is a function that has access to the variables in its outer scope, even after the outer function has finished executing.
	// closure can be used to create private variables and functions.
// if we return that value it is public if we do not return value it is private.


