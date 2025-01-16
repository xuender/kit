package cont_test

import (
	"fmt"

	"github.com/xuender/kit/v2/cont"
)

func ExampleReduce() {
	sum := cont.Reduce([]int{1, 2, 3, 4}, func(val1, val2 int) int {
		return val1 + val2
	})

	fmt.Println(sum)

	// Output:
	// 10
}
