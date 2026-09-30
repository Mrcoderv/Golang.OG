package cmd

import (
	"fmt"
	"os"
)

func DeleteFile(filename string) {
	err := os.Remove(filename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File deleted:", filename)
}