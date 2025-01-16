package cont_test

import (
	"fmt"
	"slices"

	"github.com/xuender/kit/v2/cont"
)

func ExampleDifference() {
	fmt.Println(slices.Collect(cont.Difference([]int{1, 2})))
	fmt.Println(slices.Collect(cont.Difference([]int{1, 2}, []int{2, 3})))

	for num := range cont.Difference([]int{1, 2}, []int{2, 3}) {
		fmt.Println(num)

		break
	}

	// Output:
	// [1 2]
	// [1]
	// 1
}

func ExampleDifferenceFunc() {
	equal := func(a, b int) bool { return a == b }
	fmt.Println(slices.Collect(cont.DifferenceFunc(equal, []int{1, 2})))
	fmt.Println(slices.Collect(cont.DifferenceFunc(equal, []int{1, 2}, []int{2, 3})))

	for num := range cont.DifferenceFunc(equal, []int{1, 2}, []int{2, 3}) {
		fmt.Println(num)

		break
	}

	// Output:
	// [1 2]
	// [1]
	// 1
}
