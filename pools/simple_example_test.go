package pools_test

import (
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/xuender/kit/pools"
)

// ExampleSimple is an example function.
func ExampleSimple() {
	pool := pools.NewSimple(2, func(_, _ int) {
		fmt.Println("a")
	})
	defer pool.Close()

	pool.Post(lo.Range(3)...)
	time.Sleep(time.Millisecond)

	// Output:
	// a
	// a
	// a
}
