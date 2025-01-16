package cont_test

import (
	"fmt"

	"github.com/xuender/kit/v2/cont"
)

func ExampleAll() {
	equals3 := func(num int) bool { return num == 3 }

	fmt.Println(cont.All([]int{3, 3, 3, 3}, equals3))
	fmt.Println(cont.All([]int{3, 3, 1, 3}, equals3))

	// Output:
	// true
	// false
}
