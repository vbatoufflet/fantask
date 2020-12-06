// Package fantask provides concurrent tasks control through a shared execution
// context.
package fantask

import "context"

// Tasks is a tasks controller.
type Tasks struct {
	tasks  []Task
	cancel context.CancelFunc
}

// New creates a new tasks controller instance.
func New(tasks ...Task) *Tasks {
	return &Tasks{tasks: tasks}
}

// Append appends new tasks to the controller tasks list.
func (t *Tasks) Append(tasks ...Task) *Tasks {
	t.tasks = append(t.tasks, tasks...)

	return t
}

// Cancel cancels the tasks execution context.
func (t *Tasks) Cancel() {
	if t.cancel != nil {
		t.cancel()
	}
}

// Run executes all tasks concurrently, waiting for termination.
func (t *Tasks) Run(ctx context.Context) error {
	if len(t.tasks) == 0 {
		return nil
	}

	ctx, t.cancel = context.WithCancel(ctx)
	defer t.cancel()

	errCh := make(chan error, len(t.tasks))

	for _, task := range t.tasks {
		go func(task Task) {
			errCh <- task.Run(ctx)
		}(task)
	}

	// Wait for any of the tasks to stop, then cancel context. Once done,
	// consume all remainding tasks.
	err := <-errCh

	t.cancel()

	for i, n := 1, cap(errCh); i < n; i++ {
		<-errCh
	}

	return err
}

// Task is a task interface.
type Task interface {
	Run(ctx context.Context) error
}

// TaskFunc is a task function satisfying the fantask.Task interface.
type TaskFunc func(ctx context.Context) error

// Run executes the task function.
func (t TaskFunc) Run(ctx context.Context) error {
	return t(ctx)
}
