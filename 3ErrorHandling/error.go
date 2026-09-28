// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("data.txt")  // here we caught the error returned by os.Open

// 	if err != nil {     // herfe
// 		fmt.Println("Error:", err)
// 		return

// 	}

// 	defer file.Close()

//		fmt.Println("File opened successfully")
//	}
//
// now handling the error using the function

// import (
// 	"fmt"
// 	"os"
// )

// func handleError(err error) {
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }

// func main() {
// 	file, err := os.Open("data.txt") // here we caught the error returned by os.Open

// 	if err != nil { // here
// 		handleError(err)
// 		return
// 	}
// 	file.Close()
// }

// now  handleling error using the /. using the error library
// error.new()
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	err := errors.New("this is a custom error message") // here we can create inline  error handleing
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }

// 3.  wrapup error
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	orginalErr := errors.New("34567")

// 	wrappedErr := fmt.Errorf("this is my error:%w", orginalErr)
// 	fmt.Println("Wrapped Error:", wrappedErr)
// 	unwrappedErr := errors.Unwrap(wrappedErr)  // converting the error to the original error..
// 	fmt.Println("Unwrapped Error:", unwrappedErr)
// }

// struct error
package main

import (
	"fmt"
)

type MyError struct {
	code    int
	message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("error %d: %s", e.code, e.message)
}

func main() {

	Roll := 123
	Name := ""

	fmt.Println("Roll:", Roll)
	fmt.Println("Name:", Name)

	if Roll < 0 {
		err := MyError{
			code:    133,
			message: "Roll number cannot be negative",
		}

		fmt.Println(err)
	}

	if Name == "" {
		err := MyError{
			code:    133,
			message: "Name cannot be empty",
		}

		fmt.Println(err)
	}

}
