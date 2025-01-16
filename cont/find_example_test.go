package cont_test

import (
	"fmt"

	"github.com/xuender/kit/v2/cont"
)

func ExampleFind() {
	fmt.Println(cont.Find([]int{1, 2, 3}, func(item int) bool { return item > 1 }))
	fmt.Println(cont.Find([]int{1, 2, 3}, func(item int) bool { return item > 10 }))

	// Output:
	// 2 true
	// 0 false
}
