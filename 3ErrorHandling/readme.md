visit : this https://medium.com/@raghavp791/error-handling-in-go-mastering-go-ep-003-542b5cc0e884

// error handling in go refers to the handling the unexcepted error eg built in or created erro. done by returning error values 


// go uses the function to give the error . like calling the function to throw error .

like "error" package 

// now first see the cheking method

2. using the comparision.


	if err != nil {     // herfe
		fmt.Println("Error:", err)
		return
	
	}


 1.  using the specific comparision 
     errors.Is(err, os.ErrNotExist)
2. using the type comparision
check for a specific error 
type var pathErr *os.PathError

if errors.As(err, &pathErr) {
    fmt.Println("Path error occurred")
}

...
 etc can be used for the caughting the error 
 /// ## now for the    handling the error 
 there are several method

1 . using the error handleloing function
1 .errors.New: err := errors.New("this is a custom error message") // here we can use inline  error handleing

this is use to handole the in leine error 
2. errors.Unwrap:   here we create the layer of error . it is like the abstraction
  // like we dont want to display only error we can cate
  handling unexpected errors grose them

// 
errors.is is used to check the 
// structure error on go
  like checking the error for the 


// panic and recover in go 
// panic is the way is to suddnly exectingthe  execution of the current function.
like braak 



.> 



