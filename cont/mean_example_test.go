package cont_test

import (
	"fmt"

	"github.com/xuender/kit/v2/cont"
)

func ExampleMean() {
	fmt.Println(cont.Mean([]int{1, 2, 3}))
	fmt.Println(cont.Mean([]int{}))

	// Output:
	// 2
	// 0
}
