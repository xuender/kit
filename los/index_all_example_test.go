package los_test

import (
	"fmt"

	"github.com/xuender/kit/los"
)

func ExampleIndexAll() {
	fmt.Println(los.IndexAll([]int{1, 2, 3, 4, 3, 2, 1}, 1))
	fmt.Println(los.IndexAll([]int{1, 2, 3, 4, 3, 2, 1}, 2))
	fmt.Println(los.IndexAll([]int{1, 2, 3, 4, 3, 2, 1}, 4))
	fmt.Println(los.IndexAll([]int{1, 2, 3, 4, 3, 2, 1}, 5))

	// Output:
	// [0 6]
	// [1 5]
	// [3]
	// []
}

func ExampleIndexAllFunc() {
	fmt.Println(los.IndexAllFunc([]int{1, 2, 3, 4, 3, 2, 1}, func(num int) bool { return num == 1 }))
	fmt.Println(los.IndexAllFunc([]int{1, 2, 3, 4, 3, 2, 1}, func(num int) bool { return num == 2 }))
	fmt.Println(los.IndexAllFunc([]int{1, 2, 3, 4, 3, 2, 1}, func(num int) bool { return num == 4 }))
	fmt.Println(los.IndexAllFunc([]int{1, 2, 3, 4, 3, 2, 1}, func(num int) bool { return num == 5 }))

	// Output:
	// [0 6]
	// [1 5]
	// [3]
	// []
}
