package oss_test

import (
	"fmt"

	"github.com/xuender/kit/oss"
)

// ExampleIsRelease 是否是发布模式例子.
func ExampleIsRelease() {
	fmt.Println(oss.IsRelease())
	oss.Version = "1.0.1"
	fmt.Println(oss.IsRelease())

	// Output:
	// false
	// true
}
