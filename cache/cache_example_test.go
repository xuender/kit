package cache_test

import (
	"fmt"
	"time"

	"github.com/xuender/kit/cache"
)

// ExampleNew is an example function.
func ExampleNew() {
	cac := cache.New[int, int](time.Millisecond, time.Millisecond)
	cac.SetDuration(1, 1, time.Millisecond*3)
	cac.Set(2, 1)

	fmt.Println(cac.GetNoExtension(1))
	fmt.Println(cac.GetNoExtension(3))
	fmt.Println(cac.Get(3))
	fmt.Println(cac.Len())

	time.Sleep(time.Millisecond * 2)
	fmt.Println(cac.Len())
	cac.Delete(1)
	fmt.Println(cac.Len())

	// Output:
	// 1 true
	// 0 false
	// 0 false
	// 2
	// 1
	// 0
}

// ExampleNewStringKey is an example function.
func ExampleNewStringKey() {
	cac := cache.NewStringKey[int](time.Millisecond, time.Millisecond)
	cac.SetDuration("key1", 1, time.Millisecond*3)
	cac.Set("key2", 1)

	fmt.Println(cac.Get("key2"))
	fmt.Println(cac.Len())

	time.Sleep(time.Millisecond * 2)
	fmt.Println(cac.Len())

	// Output:
	// 1 true
	// 2
	// 1
}
