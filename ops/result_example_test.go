package ops_test

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xuender/kit/v2/ops"
)

func ExampleResult() {
	fmt.Println(ops.Result(1, 2, 3, 4))

	// Output:
	// 1
}

func ExampleResult1() {
	fmt.Println(ops.Result1(1, 2, 3, 4))
	fmt.Println(ops.Result1(strconv.ParseInt("2", 10, 64)))

	// Output:
	// 1
	// 2
}

func ExampleResult2() {
	fmt.Println(ops.Result2(1, 2, 3, 4))
	fmt.Println(os.IsNotExist(ops.Result2(os.Stat("not exist"))))

	// Output:
	// 2
	// true
}

func ExampleResult3() {
	fmt.Println(ops.Result3(1, 2, 3, 4))

	// Output:
	// 3
}

func ExampleResult4() {
	fmt.Println(ops.Result4(1, 2, 3, 4))

	// Output:
	// 4
}
