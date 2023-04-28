package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/xuender/kit/cache"
	"github.com/xuender/kit/logs"
)

func test() {
	cac := cache.New[int, int](time.Second, time.Second)

	for i := 0; i < 1000; i++ {
		cac.Set(i, i)
	}

	logs.I.Println("length", cac.Len())
}

func main() {
	for i := 0; i < 10; i++ {
		test()
	}

	runtime.GC()
	fmt.Println("sleep")
	time.Sleep(time.Second)
}
