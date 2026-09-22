package main

import "fmt"

// USES 
// TO CLOSET THE FILE , DATA BASE , 
//defer schedules a function call to execute when the surrounding function returns.
// 1. To 

func main() {

    defer fmt.Println("This runs later")  //  HERO KI ENTRY LAST MAI HOTI HAI 


    fmt.Println("Hello")
    fmt.Println("Welcome")
}