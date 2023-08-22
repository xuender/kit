// nolint: dupword
package ordered_test

import (
	"fmt"

	"github.com/xuender/kit/ordered"
)

func ExampleSet_Has() {
	set := ordered.NewSet(2, 1, 3)

	fmt.Println(set.Has(2))
	fmt.Println(set.Has(1))
	fmt.Println(set.Has(3))
	fmt.Println(set.Has(0))
	fmt.Println(set.Has(4))

	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleSet_Add() {
	set := ordered.NewSet(1)

	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(2)

	fmt.Println(len(set))
	fmt.Println(set)

	// Output:
	// 3
	// [1 2 3]
}
