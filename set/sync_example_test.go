package set_test

import (
	"fmt"
	"sort"

	"github.com/xuender/kit/set"
)

// ExampleNewSync is an example function.
func ExampleNewSync() {
	nums := set.NewSync(1, 2, 3)

	fmt.Println(nums.Len())
	fmt.Println(nums.Add(3, 4, 5).Len())

	fmt.Println(nums.Has(0))
	fmt.Println(nums.Has(3))

	nums.Delete(2)
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
