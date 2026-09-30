package cmd

import (
	"fmt"
	"os"
)

func CreateFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	file.Close()
	fmt.Println("File created:", filename)
}
