package los_test

import (
	"fmt"

	"github.com/xuender/kit/los"
)

// ExampleEveryNil is an example function.
func ExampleEveryNil() {
	fmt.Println(los.EveryNil(nil, nil))
	fmt.Println(los.EveryNil(nil, 1))

	// Output:
	// true
	// false
}

// ExampleSomeNil is an example function.
func ExampleSomeNil() {
	fmt.Println(los.SomeNil(nil, 1))
	fmt.Println(los.SomeNil(1, 2))

	// Output:
	// true
	// false
}
