package fantask_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"batou.dev/fantask"
)

func Test_TasksEmpty(t *testing.T) {
	tasks := fantask.New()

	errCh := make(chan error)
	go func() { errCh <- tasks.Run(context.Background()) }()

	select {
	case err := <-errCh:
		assert.Nil(t, err)

	case <-time.After(100 * time.Millisecond):
		assert.Fail(t, "timeout reached")
	}
}

func Test_TasksOne(t *testing.T) {
	taskErr := errors.New("foo")

	tasks := fantask.New(fantask.TaskFunc(func(context.Context) error { return taskErr }))

	errCh := make(chan error)
	go func() { errCh <- tasks.Run(context.Background()) }()

	select {
	case err := <-errCh:
		assert.Equal(t, taskErr, err)

	case <-time.After(1 * time.Second):
		assert.Fail(t, "timeout reached")
	}
}

func Test_TasksMany(t *testing.T) {
	taskErr := errors.New("foo")

	tasks := fantask.New(
		fantask.TaskFunc(func(context.Context) error { return taskErr }),
	)

	tasks.Append(
		fantask.TaskFunc(func(ctx context.Context) error { <-ctx.Done(); return nil }),
	)

	errCh := make(chan error)
	go func() { errCh <- tasks.Run(context.Background()) }()

	select {
	case err := <-errCh:
		assert.Equal(t, taskErr, err)

	case <-time.After(1 * time.Second):
		assert.Fail(t, "timeout reached")
	}
}

func Test_TasksCancel(t *testing.T) {
	tasks := fantask.New(
		fantask.TaskFunc(func(ctx context.Context) error { <-ctx.Done(); return nil }),
		fantask.TaskFunc(func(ctx context.Context) error { <-ctx.Done(); return nil }),
	)

	errCh := make(chan error)
	go func() { errCh <- tasks.Run(context.Background()) }()

	time.Sleep(100 * time.Millisecond)
	tasks.Cancel()

	select {
	case err := <-errCh:
		assert.Nil(t, err)

	case <-time.After(1 * time.Second):
		assert.Fail(t, "timeout reached")
	}
}
