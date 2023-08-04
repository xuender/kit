package pools_test

import (
	"testing"
	"time"

	"github.com/xuender/kit/pools"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestNew(t *testing.T) {
	t.Parallel()

	poo := pools.New(10, func(data, num int) int {
		return data
	})
	defer poo.Close()

	poo.Post([]int{1})

	time.Sleep(time.Millisecond)
}
