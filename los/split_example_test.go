package los_test

import (
	"fmt"

	"github.com/xuender/kit/los"
)

// ExampleSplitStr is an example function.
func ExampleSplitStr() {
	fmt.Println(los.SplitStr("123,456", ','))
	fmt.Println(los.SplitStr("a-b_c", '-', '_'))

	// Output:
	// [123 456]
	// [a b c]
}

// ExampleSplit is an example function.
func ExampleSplit() {
	fmt.Println(los.Split(
		[]int{1, 0, 2, 3, 0, 4, 0, 5},
		func(num, _ int) bool { return num == 0 },
	))

	// Output:
	// [[1] [2 3] [4] [5]]
}
