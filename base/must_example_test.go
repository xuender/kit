package base_test

import (
	"fmt"

	"github.com/xuender/kit/base"
)

func ExampleMust1() {
	fmt.Println(base.Must1(1, nil))

	// Output:
	// 1
}

func ExampleMust2() {
	fmt.Println(base.Must2(1, 2, nil))

	// Output:
	// 1 2
}

func ExampleMust3() {
	fmt.Println(base.Must3(1, 2, 3, nil))

	// Output:
	// 1 2 3
}
