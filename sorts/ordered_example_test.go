package sorts_test

import (
	"fmt"

	"github.com/xuender/kit/sorts"
)

// ExampleOrdered is an example function.
func ExampleOrdered() {
	ints := []int{2, 1, 5, 4, 3}
	sorts.Ordered(ints)
	fmt.Println(ints)

	uint32s := []uint32{2, 1, 5, 4, 3}
	sorts.Ordered(uint32s)
	fmt.Println(uint32s)

	floats := []float32{2.2, 1.1, 5.5, 4.4, 3.3}
	sorts.Ordered(floats)
	fmt.Println(floats)
	// Output:
	// [1 2 3 4 5]
	// [1 2 3 4 5]
	// [1.1 2.2 3.3 4.4 5.5]
}
