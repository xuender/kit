package types_test

import (
	"fmt"

	"github.com/xuender/kit/v2/seq"
	"github.com/xuender/kit/v2/types"
)

func ExampleT() {
	fmt.Println(types.T(1, "a"))

	// Output:
	// {1 a}
}

func ExampleTuples() {
	fmt.Println(types.Tuples(seq.Range2(2)))

	// Output:
	// [{0 0} {1 1}]
}
