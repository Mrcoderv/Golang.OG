// worker pool
package main

type Task struct {
	ID int
}


func main() {
	// create a new worker pool with 5 workers
	pool := NewWorkerPool(5)

	// start the worker pool
	pool.Start()

	// submit some tasks to the worker pool
	for i := 0; i < 10; i++ {
		task := &Task{ID: i}
		pool.Submit(task)
	}
	// stop the worker pool
	pool.Stop()
}
func Task(ID int) *Task {
	return &Task{ID: ID}
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		tasks:      make(chan *Task),
	}
}
