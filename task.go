package fantask

import (
	"context"
)

// Task is a group task interface.
type Task interface {
	Run(context.Context) error
}

// TaskFunc is a group task function satisfying the task interface.
type TaskFunc func(context.Context) error

// Run executes the group task function.
func (t TaskFunc) Run(ctx context.Context) error {
	return t(ctx)
}
