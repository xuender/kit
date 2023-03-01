package base_test

import (
	"fmt"

	"github.com/xuender/kit/base"
)

func ExampleResult1() {
	fmt.Println(base.Result1(1, 2, 3, 4))

	// Output:
	// 1
}

func ExampleResult2() {
	fmt.Println(base.Result2(1, 2, 3, 4))

	// Output:
	// 2
}

func ExampleResult3() {
	fmt.Println(base.Result3(1, 2, 3, 4))

	// Output:
	// 3
}
