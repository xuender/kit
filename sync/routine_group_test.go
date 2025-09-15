package sync_test

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	ksync "github.com/xuender/kit/sync"
)

func TestNewRoutineGroup_panic(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, ksync.ErrSizeLessZero.Error(), func() {
		ksync.NewRoutineGroup(0)
	})
}

func Benchmark_lock(b *testing.B) {
	var (
		lock  sync.Mutex
		count int
	)

	b.ResetTimer()

	for range b.N {
		lock.Lock()

		count++

		lock.Unlock()
	}
}

func Benchmark_tryLock(b *testing.B) {
	var (
		lock  sync.Mutex
		count int
	)

	b.ResetTimer()

	for range b.N {
		lock.TryLock()

		count++

		lock.Unlock()
	}
}

func Benchmark_pass(b *testing.B) {
	var count atomic.Int64

	b.ResetTimer()

	for range b.N {
		count.Add(1)
	}
}

func Benchmark_ch(b *testing.B) {
	cha := make(chan int, 1)

	b.ResetTimer()

	for n := range b.N {
		cha <- n

		<-cha
	}
}
