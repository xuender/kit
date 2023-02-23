package base_test

import (
	"fmt"

	"github.com/xuender/kit/base"
)

func ExampleParam1() {
	fmt.Println(base.Param1(1, 2, 3, 4))

	// Output:
	// 1
}

func ExampleParam2() {
	fmt.Println(base.Param2(1, 2, 3, 4))

	// Output:
	// 2
}

func ExampleParam3() {
	fmt.Println(base.Param3(1, 2, 3, 4))

	// Output:
	// 3
}
