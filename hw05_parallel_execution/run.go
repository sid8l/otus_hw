package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Consumer(tasks chan Task, limitErrorExceeded chan struct{}, errorsQuantity *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		if err := task(); err != nil {
			atomic.AddInt64(errorsQuantity, -1)
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	errorsQuantity := int64(m)
	limitErrorExceeded := make(chan struct{})
	taskQueue := make(chan Task)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go Consumer(taskQueue, limitErrorExceeded, &errorsQuantity, &wg)
	}

	for i := 0; i < len(tasks); i++ {
		taskQueue <- tasks[i]
		if atomic.LoadInt64(&errorsQuantity) <= 0 {
			break
		}
	}

	close(taskQueue)
	wg.Wait()

	if errorsQuantity <= 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
