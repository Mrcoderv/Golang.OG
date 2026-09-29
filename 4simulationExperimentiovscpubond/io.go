package main

import (
	"fmt"
	"os"
)
func ioTask(id int) {

	path := "data.txt"

	data := "data " + fmt.Sprintf("task_%d", id) + "\n"

	file, err := os.OpenFile(
		path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(data)

	if err != nil {
		panic(err)
	}

	fmt.Printf("I/O task %d completed\n", id)
}
