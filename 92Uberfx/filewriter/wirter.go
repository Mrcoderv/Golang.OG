package main

import (
	"context"
	"os"

	"go.uber.org/fx"
)

type FileWriter struct {
	FileName string
	file     *os.File
}

func NewFileWriter(lc fx.Lifecycle) (*FileWriter, error) {
	file, err := os.OpenFile(
		"output.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, err
	}

	writer := &FileWriter{
		FileName: "output.txt",
		file:     file,
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return writer.Close()
		},
	})

	return writer, nil
}

func (w *FileWriter) Write(content string) error {
	_, err := w.file.WriteString(content + "\n")
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
