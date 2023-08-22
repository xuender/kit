package ordered_test

import (
	"fmt"

	"github.com/xuender/kit/ordered"
)

func ExampleInsertAes() {
	fmt.Println(ordered.InsertAes([]int{}, 1))
	fmt.Println(ordered.InsertAes([]int{3}, 1))
	fmt.Println(ordered.InsertAes([]int{3}, 5))
	fmt.Println(ordered.InsertAes([]int{1, 3}, 2))
	fmt.Println(ordered.InsertAes([]int{1, 1, 2, 3, 3, 3, 4}, 4))

	// Output:
	// [1]
	// [1 3]
	// [3 5]
	// [1 2 3]
	// [1 1 2 3 3 3 4 4]
}

func ExampleInsertDesc() {
	fmt.Println(ordered.InsertDesc([]int{}, 3))
	fmt.Println(ordered.InsertDesc([]int{1}, 3))
	fmt.Println(ordered.InsertDesc([]int{3}, 1))
	fmt.Println(ordered.InsertDesc([]int{3, 1}, 2))
	fmt.Println(ordered.InsertDesc([]int{4, 3, 3, 3, 2, 1, 1}, 4))

	// Output:
	// [3]
	// [3 1]
	// [3 1]
	// [3 2 1]
	// [4 4 3 3 3 2 1 1]
}
