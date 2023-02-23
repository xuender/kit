package base_test

import (
	"fmt"

	"github.com/xuender/kit/base"
)

func ExampleIf() {
	fmt.Println(base.If(true, 1, 2))
	fmt.Println(base.If(false, 1, 2))

	// Output:
	// 1
	// 2
}
