package times_test

import (
	"fmt"

	"github.com/xuender/kit/times"
)

func ExampleNewIDWorkerByMachine() {
	worker := times.NewIDWorkerByMachine(1, 19)
	size := 10000
	ids := make(map[int64]int, size)

	for i := range size {
		ids[worker.ID()] = i
	}

	fmt.Println(len(ids))

	// Output:
	// 10000
}

func ExampleNewIDWorker() {
	worker := times.NewIDWorker()
	size := 10000
	ids := make(map[int64]int, size)

	for i := range size {
		ids[worker.ID()] = i
	}

	fmt.Println(len(ids))

	// Output:
	// 10000
}

func ExampleNewIDWorkerByKey() {
	worker1 := times.NewIDWorkerByKey("a")
	worker2 := times.NewIDWorkerByKey("b")

	fmt.Println(worker1.ID() != worker2.ID())

	// Output:
	// true
}

func ExampleIDWorker_IDs() {
	worker := times.NewIDWorker()
	size := 10000
	ids := make(map[int64]int, size)

	for i, uid := range worker.IDs(size) {
		ids[uid] = i
	}

	fmt.Println(len(ids))

	// Output:
	// 10000
}
