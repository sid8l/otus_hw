package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func executeTask(ch <-chan Task, errorsHappened *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range ch {
		if err := task(); err != nil {
			atomic.AddInt32(errorsHappened, 1)
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	errorsExceeded := false
	var errorsHappened int32
	wg := &sync.WaitGroup{}
	ch := make(chan Task)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go executeTask(ch, &errorsHappened, wg)
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errorsHappened) >= int32(m) && m != 0 {
			errorsExceeded = true
			break
		}
		ch <- task
	}
	close(ch)
	wg.Wait()
	if errorsExceeded {
		return ErrErrorsLimitExceeded
	}
	return nil
}
