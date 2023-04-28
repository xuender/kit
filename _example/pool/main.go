package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/syncs"
)

func test() {
	pool := syncs.NewPool(3, func(elem, num int) int {
		logs.D.Println(num, elem)

		return elem
	})

	pool.Post([]int{1, 2, 3})
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
