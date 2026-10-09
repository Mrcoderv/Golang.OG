
package main

import "fmt"

type WriteService struct {
	writer *FileWriter
}

// Constructor with dependency injection  { here the WriteService depends on FileWriter, which is injected into it.}
func NewWriteService(writer *FileWriter) *WriteService {
	return &WriteService{
		writer: writer,
	}
}

func (s *WriteService) Execute() error {  // here we use the injected FileWriter to write content to the file.
	err := s.writer.Write("Hello from Uber FX IoC!")
	if err != nil {
		return err
	}

	fmt.Println("File written successfully")
	return nil
}