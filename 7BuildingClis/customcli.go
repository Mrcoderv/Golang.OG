package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: app <command>")
		return
	}

	command := os.Args[1]

	switch command {
	case "greet":
		if len(os.Args) < 3 {
			fmt.Println("Usage: app greet <name>")
			return
		}

		name := os.Args[2]
		fmt.Println("Hello,", name)

	case "version":
		fmt.Println("Version 1.0.0")

	default:
		fmt.Println("Unknown command:", command)
	}
}
