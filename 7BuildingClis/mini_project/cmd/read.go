package cmd

import (
	"fmt"
	"os"
)

func ReadFile(filename string) {
	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))
}