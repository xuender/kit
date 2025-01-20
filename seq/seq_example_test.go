package seq_test

import (
	"fmt"

	"github.com/xuender/kit/v2/seq"
	"github.com/xuender/kit/v2/types"
)

func ExampleSeq() {
	for val := range seq.Seq(1, 2, 3) {
		fmt.Println(val)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleSeq2() {
	for key, val := range seq.Seq2(types.T(1, "one"), types.T(2, "two"), types.T(3, "three")) {
		if key > 2 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 1 one
	// 2 two
}
