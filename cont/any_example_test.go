package cont_test

import (
	"fmt"

	"github.com/xuender/kit/v2/cont"
)

func ExampleAny() {
	greaterThan0 := func(num int) bool { return 0 > num }
	greaterThan2 := func(num int) bool { return 2 > num }

	fmt.Println(cont.Any([]int{1, 2}, greaterThan0))
	fmt.Println(cont.Any([]int{1, 2}, greaterThan2))

	// Output:
	// false
	// true
}
