package main

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/fx"
)

// Abstraction
type Writer interface {
	Write(content string) error
}

// Concrete implementation
type FileWriter struct {
	file *os.File
}

// Constructor
func NewFileWriter(lc fx.Lifecycle) (*FileWriter, error) {
	file, err := os.OpenFile(
		"output.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, err
	}

	writer := &FileWriter{file: file}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return writer.Close()
		},
	})

	return writer, nil
}

func (w *FileWriter) Write(content string) error {
	_, err := w.file.WriteString(content + "\n")
	fmt.Println("[Fw] Writing to    file:", content)
	return err
}

func (w *FileWriter) Close() error {
	if w.file == nil {
		return nil
	}

	err := w.file.Close()
	w.file = nil
	return err
}

// Compile-time interface check
var _ Writer = (*FileWriter)(nil)
