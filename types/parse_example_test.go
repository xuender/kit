package types_test

import (
	"fmt"

	"github.com/xuender/kit/types"
)

func ExampleFormatUint() {
	fmt.Println(types.FormatUint[uint32](123))

	// Output:
	// 123
}
