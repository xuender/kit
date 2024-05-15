package main

import (
	"log/slog"
	"time"

	"github.com/xuender/kit/sync"
)

func main() {
	var size int32 = 7

	group := sync.NewRoutineGroup(size)

	for idx := 0; idx < 100; idx++ {
		group.Incr()

		go func(gro *sync.RoutineGroup, num int) {
			slog.Info("start---------", "num", num)
			time.Sleep(time.Second)
			slog.Info("end", "num", num)
			gro.Done()
		}(group, idx)
	}

	group.Wait()
	slog.Info("done")
}
