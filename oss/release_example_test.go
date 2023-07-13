package oss_test

import (
	"fmt"
	"os"

	"github.com/xuender/kit/oss"
)

// ExampleIsRelease 是否是发布模式例子.
func ExampleIsRelease() {
	fmt.Println(oss.IsRelease())
	os.Setenv("GIN_MODE", "release")
	fmt.Println(oss.IsRelease())

	// Output:
	// false
	// true
}
