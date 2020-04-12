package fantask

import "context"

// IgnoreCanceled returns nil on context.Caneled error. Errors are returned
// unmodified otherwise.
func IgnoreCanceled(err error) error {
	if err == context.Canceled {
		return nil
	}

	return err
}
