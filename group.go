package fantask

import (
	"context"
	"os"
	"os/signal"
)

// Group is a tasks group.
type Group struct {
	tasks  []Task
	cancel context.CancelFunc
}

// NewGroup creates a new tasks group instance.
func NewGroup() *Group {
	return &Group{}
}

// Add registers a new task into the group.
func (g *Group) Add(task Task) *Group {
	g.tasks = append(g.tasks, task)
	return g
}

// Cancel cancels the group execution context.
func (g *Group) Cancel() {
	if g.cancel != nil {
		g.cancel()
	}
}

// CancelWithSignals cancels the group execution context when signals are
// received. If no signals provided, all signals will be handled.
func (g *Group) CancelWithSignals(sig ...os.Signal) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, sig...)

	<-ch

	if g.cancel != nil {
		g.cancel()
	}
}

// Run starts all registered group tasks, waiting for termination.
func (g *Group) Run(ctx context.Context) error {
	if len(g.tasks) == 0 {
		return nil
	}

	ctx, g.cancel = context.WithCancel(ctx)
	defer g.cancel()

	// Start all tasks
	errCh := make(chan error, len(g.tasks))

	for _, task := range g.tasks {
		go func(task Task) {
			errCh <- task.Run(ctx)
		}(task)
	}

	// Wait for first task to stop, then cancel context
	err := <-errCh

	g.cancel()

	// Wait for remainding tasks
	for i, n := 1, cap(errCh); i < n; i++ {
		<-errCh
	}

	return err
}
