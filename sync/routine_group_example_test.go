package sync_test

import (
	"fmt"

	"github.com/xuender/kit/sync"
)

func ExampleNewRoutineGroup() {
	group := sync.NewRoutineGroup(3)
	for i := 0; i < 5; i++ {
		group.Incr()

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

func ExampleRoutineGroup_Incr() {
	group := sync.NewRoutineGroup(3)
	for i := 0; i < 5; i++ {
		group.Incr()

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
