package times_test

import (
	"fmt"

	"github.com/xuender/kit/times"
)

// nolint: gochecknoglobals
var _time, _ = times.Parse("2023-08-02 12:11:30")

func ExampleFormat() {
	fmt.Println(times.Format(_time))

	// Output:
	// 2023-08-02 12:11:30
}

func ExampleFormatDate() {
	fmt.Println(times.FormatDate(_time))

	// Output:
	// 2023-08-02
}

func ExampleFormatTime() {
	fmt.Println(times.FormatTime(_time))

	// Output:
	// 12:11:30
}
