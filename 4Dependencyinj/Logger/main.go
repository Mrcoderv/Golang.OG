package main

import "fmt"

type Logger struct{}

func (Logger) Log(message string) {
	fmt.Println(message)
}

type UserService struct {
	logger Logger
}

type Errors struct {
	logger Logger
}

func NewErrors(logger Logger) Errors {
	return Errors{logger: logger}
}
func NewUserService(logger Logger) UserService { // constructure injection.  like transformaing the power of the logger into the user service.  this is a dependency injection
	return UserService{
		logger: logger,
	}
}

func (u UserService) CreateUser() {
	u.logger.Log("User created")
}
func (u Errors) LogError(message string) {
	u.logger.Log("Error: " + message)
}

type Info struct {
	logger Logger
}

func NewInfo(logger Logger) Info {
	return Info{logger: logger}
}
func (i Info) LogInfo(message string) {
	i.logger.Log("Info: " + message)  // here the log method is called so the logger is a dependency of the Info struct.  this is a dependency injection.  the logger is injected into the Info struct.
}
func (i Info) Users(message string) {
	i.logger.Log("Info: " + message)
}
func main() {
	logger := Logger{} // dependency creation.
	error := NewErrors(logger)
	info := NewInfo(logger)
	error.LogError("An error occurred")     // using the dependency.
	info.LogInfo("This is an info message") // using the dependency.

	info.Users("This is an info message") // using the dependency.
	userService := NewUserService(logger) // dependency injection.

	userService.CreateUser() // using the dependency.
}
