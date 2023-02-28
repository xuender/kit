// nolint: dupword
package base_test

import (
	"fmt"

	"github.com/xuender/kit/base"
)

func ExampleEveryNil() {
	fmt.Println(base.EveryNil())
	fmt.Println(base.EveryNil(nil))
	fmt.Println(base.EveryNil(nil, nil))
	fmt.Println(base.EveryNil(1))
	fmt.Println(base.EveryNil(1, nil))

	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleSomeNil() {
	fmt.Println(base.SomeNil())
	fmt.Println(base.SomeNil(nil))
	fmt.Println(base.SomeNil(nil, nil))
	fmt.Println(base.SomeNil(1))
	fmt.Println(base.SomeNil(1, nil))

	// Output:
	// true
	// true
	// true
	// false
	// true
}
