package seq_test

import (
	"fmt"

	"github.com/xuender/kit/v2/seq"
)

func ExampleZip() {
	for key, val := range seq.Zip(seq.Range(5), seq.Range(6, 20)) {
		if key > 2 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 0 6
	// 1 7
	// 2 8
}
