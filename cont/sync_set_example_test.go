package cont_test

import (
	"fmt"
	"sort"

	"github.com/xuender/kit/v2/cont"
)

func ExampleNewSyncSet() {
	nums := cont.NewSyncSet(1, 2, 3)

	fmt.Println(nums.Len())
	fmt.Println(nums.Add(3, 4, 5).Len())

	fmt.Println(nums.Has(0))
	fmt.Println(nums.Has(3))

	nums.Del(2)
	ints := nums.Slice()
	sort.Ints(ints)

	fmt.Println(ints)

	for range nums.Values() {
		fmt.Println(1)
	}

	// Output:
	// 3
	// 5
	// false
	// true
	// [1 3 4 5]
	// 1
	// 1
	// 1
	// 1
}
