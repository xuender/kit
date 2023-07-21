package counter_test

import (
	"fmt"

	"github.com/xuender/kit/counter"
)

func ExampleNumHLLPP() {
	hll := counter.NewNumHLLPP[int]()
	hll.Add(3)
	hll.Add(3)
	hll.Add(1)
	hll.Add(3)

	fmt.Println(hll.Count())

	// Output:
	// 2
}
