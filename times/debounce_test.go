package times_test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/times"
)

func TestDebounce(t *testing.T) {
	t.Parallel()

	var count int32

	ass := assert.New(t)

	play, wait := times.Debounce(func() {
		atomic.AddInt32(&count, 1)
	}, time.Millisecond*100)

	for range 10 {
		go play()
	}

	time.Sleep(time.Millisecond * 50)

	wait()

	ass.Equal(int32(1), atomic.LoadInt32(&count))
}

func TestDebounce_two(t *testing.T) {
	t.Parallel()

	var count int32

	ass := assert.New(t)

	play, wait := times.Debounce(func() {
		atomic.AddInt32(&count, 1)
	}, time.Millisecond*100)

	for range 10 {
		play()
	}

	wait()

	for range 10 {
		play()
	}

	wait()

	ass.Equal(int32(2), atomic.LoadInt32(&count))
}
