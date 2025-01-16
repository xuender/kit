package cont_test

import (
	"fmt"

	"github.com/xuender/kit/v2/cont"
)

func ExampleLastIndex() {
	fmt.Println(cont.LastIndex([]int{1, 2, 3}, 3))
	fmt.Println(cont.LastIndex([]int{1, 2, 3}, 2))
	fmt.Println(cont.LastIndex([]int{1, 2, 3}, 1))
	fmt.Println(cont.LastIndex([]int{1, 2, 3}, 4))

	// Output:
	// 2
	// 1
	// 0
	// -1
}

func ExampleLastIndexFunc() {
	fmt.Println(cont.LastIndexFunc([]int{1, 2, 3}, func(num int) bool { return num == 2 }))
	fmt.Println(cont.LastIndexFunc([]int{1, 2, 3}, func(num int) bool { return num > 10 }))

	// Output:
	// 1
	// -1
}
