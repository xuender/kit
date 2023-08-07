package set_test

import (
	"fmt"
	"sort"

	"github.com/xuender/kit/set"
)

// ExampleNewSet is an example function.
func ExampleNewSet() {
	nums := set.NewSet(1, 2, 3)

	fmt.Println(len(nums))
	fmt.Println(len(nums.Add(3, 4, 5)))

	fmt.Println(nums.Has(0))
	fmt.Println(nums.Has(3))

	delete(nums, 2)
	ints := nums.Slice()
	sort.Ints(ints)

	fmt.Println(ints)

	// Output:
	// 3
	// 5
	// false
	// true
	// [1 3 4 5]
}
