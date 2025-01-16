package cont_test

import (
	"fmt"
	"slices"

	"github.com/xuender/kit/v2/cont"
)

func ExampleIntersection() {
	fmt.Println(slices.Collect(cont.Intersection[[]int]()))
	fmt.Println(slices.Collect(cont.Intersection([]int{1, 2, 3})))
	fmt.Println(slices.Collect(cont.Intersection([]int{1, 2, 2, 3}, []int{2, 3, 4})))

	for num := range cont.Intersection([]int{1, 1, 2, 3}, []int{2, 3, 4}) {
		fmt.Println(num)

		break
	}

	// Output:
	// []
	// [1 2 3]
	// [2 3]
	// 2
}

func ExampleIntersectionFunc() {
	equal := func(a, b int) bool { return a == b }
	fmt.Println(slices.Collect(cont.IntersectionFunc[[]int](equal)))
	fmt.Println(slices.Collect(cont.IntersectionFunc(equal, []int{1, 2, 3})))
	fmt.Println(slices.Collect(cont.IntersectionFunc(equal, []int{1, 2, 2, 3}, []int{2, 3, 4})))

	for num := range cont.IntersectionFunc(equal, []int{1, 1, 2, 3}, []int{2, 3, 4}) {
		fmt.Println(num)

		break
	}

	// Output:
	// []
	// [1 2 3]
	// [2 3]
	// 2
}
