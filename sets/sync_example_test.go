package sets_test

import (
	"fmt"
	"sort"

	"github.com/xuender/kit/sets"
)

// ExampleNewSync is an example function.
func ExampleNewSync() {
	set := sets.NewSync(1, 2, 3)

	fmt.Println(set.Len())
	fmt.Println(set.Add(3, 4, 5).Len())

	fmt.Println(set.Has(0))
	fmt.Println(set.Has(3))

	set.Delete(2)
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
