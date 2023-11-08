package oss_test

import (
	"fmt"

	"github.com/xuender/kit/oss"
)

func ExampleProcInfo_String() {
	pro := oss.NewProcInfo()
	fmt.Println(pro.String() != "")

	// Output:
	// true
}
