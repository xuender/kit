package pools_test

import (
	"fmt"

	"github.com/xuender/kit/pools"
)

type data struct {
	data int
}

// ExampleNewSyncPool is an example function.
func ExampleNewSyncPool() {
	pool := pools.NewSyncPool(func() *data { return &data{1} }, func(d *data) { d.data = 1 })
	d1 := pool.Get()
	d1.data = 3
	pool.Put(d1)

	d2 := pool.Get()
	fmt.Println(d2.data)

	// Output:
	// 1
}
