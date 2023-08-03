package times_test

import (
	"fmt"
	"time"

	"github.com/xuender/kit/times"
)

func ExampleDuration_String() {
	fmt.Println(times.Duration(time.Microsecond * 3))
	fmt.Println(times.Duration(time.Hour*48 + 1))

	// Output:
	// 3微秒
	// 2天1纳秒
}

func ExampleDuration_Short() {
	fmt.Println(times.Duration(time.Hour * 30).Short())
	fmt.Println(times.Duration(time.Hour * 3).Short())
	fmt.Println(times.Duration(time.Minute * 3).Short())
	fmt.Println(times.Duration(time.Millisecond * 3).Short())
	fmt.Println(times.Duration(time.Second * 3).Short())
	fmt.Println(times.Duration(time.Second*3 + time.Hour*3 + 3).Short())

	// Output:
	// 1天6小时
	// 3小时
	// 3分钟
	// 3毫秒
	// 3秒钟
	// 3小时3秒钟
}
