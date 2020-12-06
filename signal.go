package fantask

import (
	"os"
	"os/signal"
)

// CancelWithSignal cancels the tasks execution context when any of the given
// signals is received. If no signals provided, all signals will be handled.
func CancelWithSignal(tasks *Tasks, sigs ...os.Signal) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, sigs...)

	<-ch

	tasks.Cancel()
}
