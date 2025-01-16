package cont_test

import (
	"fmt"

	"github.com/xuender/kit/v2/cont"
)

func ExampleFilter() {
	items := cont.Filter([]int{1, 2, 3, 4}, func(num int) bool { return num%2 == 0 })

	for num := range items {
		fmt.Println(num)

		break
	}

	// Output:
	// 2
}
