package pools_test

import (
	"fmt"

	"github.com/xuender/kit/pools"
)

// ExampleArrange is an example function.
func ExampleArrange() {
	fmt.Println(pools.Arrange(
		[]int{1, 1, 2, 2, 2, 3},
		func(num int) int { return num }),
	)

	// Output:
	// [2 1 3 2 1 2]
}
