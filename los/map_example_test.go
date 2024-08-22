package los_test

import (
	"fmt"

	"github.com/xuender/kit/los"
)

func ExampleMap() {
	fmt.Println(los.Map([]int{1, 2, 3, 4}, func(item int) int { return item * 2 }))

	// Output:
	// [2 4 6 8]
}
