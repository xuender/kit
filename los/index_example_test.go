package los_test

import (
	"fmt"

	"github.com/xuender/kit/los"
)

func ExampleIndexOf() {
	slice := []int{1, 2, 1, 2}

	fmt.Println(los.IndexOf(slice, []int{2, 1}))
	fmt.Println(los.IndexOf(slice, []int{1, 2, 1}))
	fmt.Println(los.IndexOf(slice, []int{2, 1, 1}))
	fmt.Println(los.IndexOf(slice, []int{1, 2, 1, 2, 1}))
	fmt.Println(los.IndexOf(slice, []int{1, 2}))

	// Output:
	// 1
	// 0
	// -1
	// -1
	// 0
}

func ExampleLastIndexOf() {
	slice := []int{1, 2, 1, 2}

	fmt.Println(los.LastIndexOf(slice, []int{2, 1}))
	fmt.Println(los.LastIndexOf(slice, []int{1, 2, 1}))
	fmt.Println(los.LastIndexOf(slice, []int{2, 1, 1}))
	fmt.Println(los.LastIndexOf(slice, []int{1, 2, 1, 2, 1}))
	fmt.Println(los.LastIndexOf(slice, []int{1, 2}))

	// Output:
	// 1
	// 0
	// -1
	// -1
	// 2
}
