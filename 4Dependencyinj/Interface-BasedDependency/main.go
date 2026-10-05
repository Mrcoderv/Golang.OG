//interface based dependecy making
package main

import "fmt"

// Interface
type Logger interface {
	Log(message string)  // this is the behaviour of the logger.  
}

// Implementation
type ConsoleLogger struct{}

func (ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

// Component
type UserService struct {
	logger Logger 
}


func NewUserService(logger Logger) UserService {
	return UserService{
		logger: logger,
	}
}

func (u UserService) CreateUser() {
	u.logger.Log("User created")
}

func main() {
	logger := ConsoleLogger{}
	

	userService := NewUserService(logger)

	userService.CreateUser()
}