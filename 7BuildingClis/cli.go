// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Usage: app <command>")
// 		return
// 	}

// 	command := os.Args[1]

// 	switch command { // by this cases we can create the command line application which can take the command from the user and perform the action based on the command provided by the user.
// 	case "hello":
// 		fmt.Println("hello!")

// 	case "version":
// 		fmt.Println("version 1.0.0")

// 	case "help":
// 		fmt.Println("available commands")
// 		fmt.Println("  hello") //
// 		fmt.Println("  version")
// 		fmt.Println("  help")

// 	default:
// 		fmt.Println("Unknown command:", command)
// 	}
// }
