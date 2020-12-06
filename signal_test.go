package fantask_test

import (
	"context"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"batou.dev/fantask"
)

func Test_CancelWithSignals(t *testing.T) {
	tasks := fantask.New(
		fantask.TaskFunc(func(ctx context.Context) error { <-ctx.Done(); return nil }),
		fantask.TaskFunc(func(ctx context.Context) error { <-ctx.Done(); return nil }),
	)

	errCh := make(chan error)

	go func() { errCh <- tasks.Run(context.Background()) }()

	go fantask.CancelWithSignal(tasks, syscall.SIGUSR1)

	time.Sleep(100 * time.Millisecond)

	err := syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)
	assert.Nil(t, err)

	select {
	case err := <-errCh:
		assert.Nil(t, err)

	case <-time.After(1 * time.Second):
		assert.Fail(t, "timeout reached")
	}
}
