package los_test

import (
	"fmt"

	"github.com/xuender/kit/los"
)

// ExamplePull is an example function.
func ExamplePull() {
	fmt.Println(los.Pull([]int{1, 2, 2, 3}, 0, 2))

	// Output:
	// [1 3]
}

// ExamplePullAt is an example function.
func ExamplePullAt() {
	fmt.Println(los.PullAt([]rune{'1', '2', '2', '3'}, 2, 3))

	// Output:
	// [49 50]
}

// ExampleRemove is an example function.
func ExampleRemove() {
	fmt.Println(los.Remove([]int{1, 2, 2, 3, 4}, func(item, _ int) bool { return item%2 == 0 }))

	// Output:
	// [1 3]
}
