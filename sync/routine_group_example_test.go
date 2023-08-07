package sync_test

import (
	"fmt"

	"github.com/xuender/kit/sync"
)

func ExampleNewRoutineGroup() {
	group := sync.NewRoutineGroup(3)
	for i := 0; i < 5; i++ {
		group.Add(1)

		go func() {
			defer group.Done()

			fmt.Println("x")
		}()
	}

	group.Wait()

	// Output:
	// x
	// x
	// x
	// x
	// x
}

func ExampleRoutineGroup_Inc() {
	group := sync.NewRoutineGroup(3)
	for i := 0; i < 5; i++ {
		group.Inc()

		go func() {
			defer group.Done()

			fmt.Println("x")
		}()
	}

	group.Wait()

	// Output:
	// x
	// x
	// x
	// x
	// x
}
