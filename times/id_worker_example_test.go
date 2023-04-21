package times_test

import (
	"fmt"

	"github.com/xuender/kit/times"
)

// ExampleNewIDWorkerByMachine is an example function.
func ExampleNewIDWorkerByMachine() {
	worker := times.NewIDWorkerByMachine(1, 19)
	size := 10000
	ids := make(map[int64]int, size)

	for i := 0; i < size; i++ {
		ids[worker.ID()] = i
	}

	fmt.Println(len(ids))

	// Output:
	// 10000
}

// ExampleNewIDWorker is an example function.
func ExampleNewIDWorker() {
	worker := times.NewIDWorker()
	size := 10000
	ids := make(map[int64]int, size)

	for i := 0; i < size; i++ {
		ids[worker.ID()] = i
	}

	fmt.Println(len(ids))

	// Output:
	// 10000
}

// ExampleNewIDWorkerByKey is an example function.
func ExampleNewIDWorkerByKey() {
	worker1 := times.NewIDWorkerByKey("a")
	worker2 := times.NewIDWorkerByKey("b")

	fmt.Println(worker1.ID() != worker2.ID())

	// Output:
	// true
}
