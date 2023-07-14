package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/pools"
)

func test() {
	pool := pools.NewSimple(3, func(elem, num int) {
		logs.D.Println(num, elem)
	})

	pool.Post(1, 2, 3)
}

func main() {
	for i := 0; i < 10; i++ {
		test()
	}

	fmt.Println("sleep")
	time.Sleep(time.Second)
	runtime.GC()
	time.Sleep(time.Second)
}
