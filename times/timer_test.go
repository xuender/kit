package times_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/times"
)

func TestWithTimerCancel(t *testing.T) {
	t.Parallel()

	cancel := times.WithTimer(time.Second, func() {
		assert.Nil(t, "error")
	})
	cancel()
	time.Sleep(time.Millisecond)
}

func TestWithTimer(t *testing.T) {
	t.Parallel()

	_ = times.WithTimer(time.Microsecond, func() {})
	time.Sleep(time.Millisecond)
}

func TestWithContextTimerCancel(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())

	times.WithContextTimer(ctx, time.Second, func() {
		assert.Nil(t, "error")
	})
	cancel()
	time.Sleep(time.Millisecond)
}

func TestWithContextTimer(t *testing.T) {
	t.Parallel()

	times.WithContextTimer(context.Background(), time.Microsecond, func() {})
	time.Sleep(time.Millisecond)
}
