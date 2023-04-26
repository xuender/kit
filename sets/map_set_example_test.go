package sets_test

import (
	"fmt"
	"sort"

	"github.com/xuender/kit/sets"
)

// ExampleNewMapSet is an example function.
func ExampleNewMapSet() {
	set := sets.NewMapSet(1, 2, 3)

	fmt.Println(len(set))
	fmt.Println(len(set.Add(3, 4, 5)))

	fmt.Println(set.Has(0))
	fmt.Println(set.Has(3))

	delete(set, 2)
	ints := set.Slice()
	sort.Ints(ints)

	fmt.Println(ints)

	// Output:
	// 3
	// 5
	// false
	// true
	// [1 3 4 5]
}
