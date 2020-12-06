package fantask

import (
	"context"
	"errors"
)

// IgnoreCanceled returns nil on context.Canceled error. Errors are returned
// unmodified otherwise.
func IgnoreCanceled(err error) error {
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}
