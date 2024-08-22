package los_test

import (
	"fmt"

	"github.com/xuender/kit/los"
)

func ExampleRange() {
	fmt.Println(los.Range[uint32](3))

	// Output:
	// [0 1 2]
}
