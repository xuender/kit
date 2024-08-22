package los_test

import (
	"fmt"

	"github.com/xuender/kit/los"
)

func ExampleFilter() {
	fmt.Println(los.Filter([]int{1, 2, 3, 4}, func(num int) bool { return num%2 == 0 }))

	// Output:
	// [2 4]
}
