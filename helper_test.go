package fantask

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IgnoreCanceled(t *testing.T) {
	err := errors.New("test error")
	assert.Error(t, err, IgnoreCanceled(err))
	assert.Nil(t, IgnoreCanceled(context.Canceled))
	assert.Nil(t, IgnoreCanceled(nil))
}
