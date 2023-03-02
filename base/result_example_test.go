package base_test

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xuender/kit/base"
)

func ExampleResult1() {
	fmt.Println(base.Result1(1, 2, 3, 4))
	fmt.Println(base.Result1(strconv.ParseInt("2", 10, 64)))

	// Output:
	// 1
	// 2
}

func ExampleResult2() {
	fmt.Println(base.Result2(1, 2, 3, 4))
	fmt.Println(os.IsNotExist(base.Result2(os.Stat("not exist"))))

	// Output:
	// 2
	// true
}

func ExampleResult3() {
	fmt.Println(base.Result3(1, 2, 3, 4))

	// Output:
	// 3
}

func ExampleResult4() {
	fmt.Println(base.Result4(1, 2, 3, 4))

	// Output:
	// 4
}
