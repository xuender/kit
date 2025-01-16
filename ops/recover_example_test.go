// nolint: nonamedreturns
package ops_test

import (
	"fmt"
	"io"

	"github.com/xuender/kit/v2/ops"
)

func callNum() (res error) {
	defer ops.Recover(func(err error) {
		res = err
	})

	panic(3)
}

func callString() (res error) {
	defer ops.Recover(func(err error) {
		res = err
	})

	panic("str")
}

func callError() (res error) {
	defer ops.Recover(func(err error) {
		res = err
	})

	panic(io.EOF)
}

func callNoRecover() (res error) {
	defer ops.Recover(func(err error) {
		res = err
	})

	return
}

func ExampleRecover() {
	fmt.Println(callNoRecover())
	fmt.Println(callString())
	fmt.Println(callError())
	fmt.Println(callNum())

	// Output:
	// <nil>
	// str
	// EOF
	// 3
}
