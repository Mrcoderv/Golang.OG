package main

import (
	"os"
)

type FileWriter struct {
	FileName string
}

// Constructor
func NewFileWriter() *FileWriter {  // function to create a new instance of FileWriter
	return &FileWriter{
		FileName: "output.txt",
	}
}

func (w *FileWriter) Write(content string) error {  // function to write content to the file
	return os.WriteFile(
		w.FileName,
		[]byte(content),
		0644,
	)
}
