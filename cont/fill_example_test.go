package cont_test

import (
	"fmt"

	"github.com/xuender/kit/v2/cont"
)

func ExampleFill() {
	items := []int{}
	cont.Fill(items, 0, 8)
	fmt.Println(items)

	items = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	cont.Fill(items, 0, 8)
	fmt.Println(items)
	cont.Fill(items, 0, 3, 5)
	fmt.Println(items)
	cont.Fill(items, 0)
	fmt.Println(items)

	// Output:
	// []
	// [1 2 3 4 5 6 7 8 0 0]
	// [1 2 3 0 0 6 7 8 0 0]
	// [0 0 0 0 0 0 0 0 0 0]
}
