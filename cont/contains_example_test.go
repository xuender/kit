package cont_test

import (
	"fmt"

	"github.com/xuender/kit/v2/cont"
)

func ExampleContains() {
	equal := func(a, b int) bool { return a == b }

	fmt.Println(cont.Contains(equal, []int{}, 2))
	fmt.Println(cont.Contains(equal, []int{1, 2, 3}, 2))
	fmt.Println(cont.Contains(equal, []int{1, 2, 3}, 4))

	// Output:
	// false
	// true
	// false
}
