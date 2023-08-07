package ordered_test

import (
	"fmt"

	"github.com/xuender/kit/ordered"
)

func ExampleSort() {
	ints := []int{2, 1, 5, 4, 3}
	ordered.Sort(ints)
	fmt.Println(ints)

	uint32s := []uint32{2, 1, 5, 4, 3}
	ordered.Sort(uint32s)
	fmt.Println(uint32s)

	floats := []float32{2.2, 1.1, 5.5, 4.4, 3.3}
	ordered.Sort(floats)
	fmt.Println(floats)
	// Output:
	// [1 2 3 4 5]
	// [1 2 3 4 5]
	// [1.1 2.2 3.3 4.4 5.5]
}
