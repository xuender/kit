package cont_test

import (
	"fmt"
	"slices"

	"github.com/xuender/kit/v2/cont"
)

func ExampleUnion() {
	fmt.Println(slices.Collect(cont.Union[[]int]()))
	fmt.Println(slices.Collect(cont.Union([]int{1, 1})))
	fmt.Println(slices.Collect(cont.Union([]int{1, 2}, []int{2, 3})))

	for num := range cont.Union([]int{1, 2}, []int{2, 3}) {
		fmt.Println(num)

		break
	}

	// Output:
	// []
	// [1]
	// [1 2 3]
	// 1
}

func ExampleUnionFunc() {
	equal := func(a, b int) bool { return a == b }

	fmt.Println(slices.Collect(cont.UnionFunc[[]int](equal)))
	fmt.Println(slices.Collect(cont.UnionFunc(equal, []int{1, 1})))
	fmt.Println(slices.Collect(cont.UnionFunc(equal, []int{1, 2}, []int{2, 3})))

	for num := range cont.UnionFunc(equal, []int{1, 2}, []int{2, 3}) {
		fmt.Println(num)

		break
	}

	// Output:
	// []
	// [1]
	// [1 2 3]
	// 1
}
