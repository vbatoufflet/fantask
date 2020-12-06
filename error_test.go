package fantask_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"batou.dev/fantask"
)

func Test_IgnoreCanceled(t *testing.T) {
	err := errors.New("test error")
	assert.Error(t, err, fantask.IgnoreCanceled(err))
	assert.Nil(t, fantask.IgnoreCanceled(context.Canceled))
	assert.Nil(t, fantask.IgnoreCanceled(nil))
}
