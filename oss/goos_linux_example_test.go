//go:build linux

package oss_test

import (
	"fmt"

	"github.com/xuender/kit/oss"
)

func ExampleIsDarwin() {
	fmt.Println(oss.IsDarwin())

	// Output:
	// false
}

func ExampleIsLinux() {
	fmt.Println(oss.IsLinux())

	// Output:
	// true
}

func ExampleIsWindows() {
	fmt.Println(oss.IsWindows())

	// Output:
	// false
}
