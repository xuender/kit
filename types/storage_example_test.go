package types_test

import (
	"fmt"

	"github.com/xuender/kit/types"
)

func ExampleStorage() {
	fmt.Println(types.Storage(1))
	fmt.Println(types.Storage(1024))
	fmt.Println(types.Storage(1024 * 1024))
	fmt.Println(types.Storage(1024 * 1024 * 1024))
	fmt.Println(types.Storage(1024 * 1024 * 1024 * 1024))
	fmt.Println(types.Storage(1024 * 1024 * 1024 * 1024 * 1024))
	fmt.Println(types.Storage(100000))
	fmt.Println(types.Storage(10000000))

	// Output:
	// 1B
	// 1K
	// 1M
	// 1G
	// 1T
	// 1P
	// 97K
	// 9M
}
