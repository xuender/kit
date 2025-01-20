package seq_test

import (
	"fmt"
	"slices"

	"github.com/xuender/kit/v2/seq"
)

func ExampleSum() {
	fmt.Println(seq.Sum(seq.Range(101)))
	fmt.Println(seq.Sum(slices.Values([]string{"a", "b", "c"})))

	// Output:
	// 5050
	// abc
}
