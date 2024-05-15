package main

import (
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/xuender/kit/sync"
)

func main() {
	var (
		size  int32 = 200
		count atomic.Int32
	)

	group := sync.NewRoutineGroup(size)

	for idx := 0; idx < 10000; idx++ {
		group.Incr()

		go func(gro *sync.RoutineGroup, num int) {
			time.Sleep(time.Millisecond * time.Duration(size))
			count.Add(1)
			gro.Done()
		}(group, idx)
	}

	group.Wait()
	slog.Info("done", "count", count.Load())
}
