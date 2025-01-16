package cont_test

import (
	"fmt"

	"github.com/xuender/kit/v2/cont"
)

func ExampleNone() {
	equals3 := func(num int) bool { return num == 3 }

	fmt.Println(cont.None([]int{1, 2, 4, 5}, equals3))
	fmt.Println(cont.None([]int{3, 3, 1, 3}, equals3))

	// Output:
	// true
	// false
}
