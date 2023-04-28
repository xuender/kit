package pools_test

import (
	"runtime"
	"testing"
	"time"

	"github.com/xuender/kit/pools"
)

func TestNew(t *testing.T) {
	t.Parallel()

	for i := 0; i < 10; i++ {
		poo := pools.New(10, func(data, num int) int {
			return data
		})
		poo.Post([]int{1})
	}

	runtime.GC()
	time.Sleep(time.Millisecond)
}
