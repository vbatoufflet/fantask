package fantask

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GroupEmpty(t *testing.T) {
	g := NewGroup()

	errCh := make(chan error)

	go func() { errCh <- g.Run(context.Background()) }()

	select {
	case err := <-errCh:
		assert.Nil(t, err)

	case <-time.After(100 * time.Millisecond):
		assert.Fail(t, "timeout reached")
	}
}

func Test_GroupOne(t *testing.T) {
	taskErr := errors.New("foo")

	g := NewGroup().
		Add(TaskFunc(func(context.Context) error { return taskErr }))

	errCh := make(chan error)

	go func() { errCh <- g.Run(context.Background()) }()

	select {
	case err := <-errCh:
		assert.Equal(t, taskErr, err)

	case <-time.After(1 * time.Second):
		assert.Fail(t, "timeout reached")
	}
}

func Test_GroupMany(t *testing.T) {
	taskErr := errors.New("foo")

	g := NewGroup().
		Add(TaskFunc(func(context.Context) error { return taskErr })).
		Add(TaskFunc(func(ctx context.Context) error { <-ctx.Done(); return nil }))

	errCh := make(chan error)

	go func() { errCh <- g.Run(context.Background()) }()

	select {
	case err := <-errCh:
		assert.Equal(t, taskErr, err)

	case <-time.After(1 * time.Second):
		assert.Fail(t, "timeout reached")
	}
}

func Test_GroupCancel(t *testing.T) {
	g := NewGroup().
		Add(TaskFunc(func(ctx context.Context) error { <-ctx.Done(); return nil })).
		Add(TaskFunc(func(ctx context.Context) error { <-ctx.Done(); return nil }))

	errCh := make(chan error)

	go func() { errCh <- g.Run(context.Background()) }()

	time.Sleep(100 * time.Millisecond)
	g.Cancel()

	select {
	case err := <-errCh:
		assert.Nil(t, err)

	case <-time.After(1 * time.Second):
		assert.Fail(t, "timeout reached")
	}
}
