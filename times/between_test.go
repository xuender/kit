package times_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/times"
)

func TestSleep(t *testing.T) {
	t.Parallel()

	assert.Greater(t, times.Sleep(1502), time.Second)
	assert.Greater(t, times.Sleep(502), time.Second)
}
