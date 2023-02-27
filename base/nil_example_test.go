// nolint: dupword
package base_test

import (
	"fmt"

	"github.com/xuender/kit/base"
)

func ExampleAllNil() {
	fmt.Println(base.AllNil())
	fmt.Println(base.AllNil(nil))
	fmt.Println(base.AllNil(nil, nil))
	fmt.Println(base.AllNil(1))
	fmt.Println(base.AllNil(1, nil))

	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleAnyNil() {
	fmt.Println(base.AnyNil())
	fmt.Println(base.AnyNil(nil))
	fmt.Println(base.AnyNil(nil, nil))
	fmt.Println(base.AnyNil(1))
	fmt.Println(base.AnyNil(1, nil))

	// Output:
	// true
	// true
	// true
	// false
	// true
}
