package main

import (
	"errors"
	"strings"
)

var ErrEmptyTitle = errors.New("title is required")

type TaskService struct {
	repository *TaskRepository
}

func NewTaskService(repository *TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (s *TaskService) List() []Task {
	return s.repository.List()
}

func (s *TaskService) Get(id int) (Task, error) {
	return s.repository.Get(id)
}

func (s *TaskService) Create(title string) (Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Task{}, ErrEmptyTitle
	}
	return s.repository.Create(title), nil
}

func (s *TaskService) Update(id int, title string, completed bool) (Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Task{}, ErrEmptyTitle
	}
	return s.repository.Update(id, title, completed)
}

func (s *TaskService) Delete(id int) error {
	return s.repository.Delete(id)
}
