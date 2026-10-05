// single responsibility principle
package main

import "fmt"

type User interface {
	GetName() string
}

type user struct {
	Name string
}

func (u user) GetName() string {
	return u.Name
}

type UserRepository struct {
	users []User
}

func (r *UserRepository) Add(user User) {
	r.users = append(r.users, user)
}

func (r *UserRepository) GetAll() []User {
	return r.users
}

type UserService struct {
	repo *UserRepository
}

func (s *UserService) PrintAllUsers() {
	users := s.repo.GetAll()
	for _, user := range users {
		fmt.Println(user.GetName())
	}
}

func main() {
	repo := &UserRepository{}
	service := &UserService{repo: repo}

	user1 := user{Name: "Alice"}
	user2 := user{Name: "Bob"}

	repo.Add(user1)
	repo.Add(user2)

	service.PrintAllUsers()
}
