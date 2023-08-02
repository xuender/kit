package los_test

import (
	"fmt"

	"github.com/xuender/kit/los"
)

// ExampleArrange is an example function.
func ExampleArrange() {
	fmt.Println(los.Arrange(
		[]int{1, 1, 2, 2, 2, 3},
		func(num int) int { return num }),
	)

	// Output:
	// [2 1 3 2 1 2]
}
