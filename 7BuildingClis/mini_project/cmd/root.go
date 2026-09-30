package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mycli <command> <filename>")
		return
	}

	command := os.Args[1]

	switch command {

	case "create":
		if len(os.Args) < 3 {
			fmt.Println("Usage: mycli create <filename>")
			return
		}

		CreateFile(os.Args[2])

	case "read":
		if len(os.Args) < 3 {
			fmt.Println("Usage: mycli read <filename>")
			return
		}

		ReadFile(os.Args[2])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: mycli delete <filename>")
			return
		}

		DeleteFile(os.Args[2])

	default:
		fmt.Println("Unknown command:", command)
	}
}
