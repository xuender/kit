package sync_test

import (
	"fmt"

	"github.com/xuender/kit/sync"
)

func ExampleNewMap() {
	smap := sync.NewMap[int, int]()

	smap.Store(1, 1)
	fmt.Println(smap.Load(1))
	fmt.Println(smap.LoadOrCreate(2, func() int { return 2 }))

	// Output:
	// 1 true
	// 2 false
}
