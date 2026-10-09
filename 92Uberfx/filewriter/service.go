package main

type WriteService struct {
	writer Writer
}

// Constructor injection
func NewWriteService(writer Writer) *WriteService {
	return &WriteService{
		writer: writer,
	}
}

func (s *WriteService) Execute() error {
	return s.writer.Write("Hello from Uber FX with interfaces!")
}
