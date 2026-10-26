package main

import (
	"errors"
	"sync"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository struct {
	mu     sync.RWMutex
	nextID int
	tasks  map[int]Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		nextID: 1,
		tasks:  make(map[int]Task),
	}
}

func (r *TaskRepository) List() []Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (r *TaskRepository) Get(id int) (Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return task, nil
}

func (r *TaskRepository) Create(title string) Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	task := Task{ID: r.nextID, Title: title}
	r.tasks[task.ID] = task
	r.nextID++
	return task
}

func (r *TaskRepository) Update(id int, title string, completed bool) (Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return Task{}, ErrTaskNotFound
	}

	task := Task{ID: id, Title: title, Completed: completed}
	r.tasks[id] = task
	return task, nil
}

func (r *TaskRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return ErrTaskNotFound
	}
	delete(r.tasks, id)
	return nil
}
