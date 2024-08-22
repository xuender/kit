package los_test

import (
	"fmt"

	"github.com/xuender/kit/los"
)

func ExampleDelete() {
	fmt.Println(los.Delete([]int{0, 1, 2, 3, 4}))
	fmt.Println(los.Delete([]int{0, 1, 2, 3, 4}, 1))
	fmt.Println(los.Delete([]int{0, 1, 2, 3, 4}, 3, 1))
	fmt.Println(los.Delete([]int{0, 1, 2, 3, 4}, 3, 4))
	fmt.Println(los.Delete([]int{0, 1, 2, 3, 4}, 2, 3))
	fmt.Println(los.Delete([]int{0, 1, 2, 3, 4}, 0, 1))
	fmt.Println(los.Delete([]int{0, 1, 2, 3, 4}, 0, 1, 2, 3, 4))

	// Output:
	// [0 1 2 3 4]
	// [0 2 3 4]
	// [0 2 4]
	// [0 1 2]
	// [0 1 4]
	// [2 3 4]
	// []
}

func ExampleDeleteBy() {
	fmt.Println(los.DeleteBy([]int{0, 1, 2, 3, 4}, 1))

	// Output:
	// [0 2 3 4]
}

func ExampleDeleteFunc() {
	fmt.Println(los.DeleteFunc([]int{0, 1, 2, 3, 4}, func(num int) bool { return num == 1 }))

	// Output:
	// [0 2 3 4]
}
