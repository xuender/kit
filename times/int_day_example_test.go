package times_test

import (
	"fmt"

	"github.com/xuender/kit/times"
)

func ExampleNow2IntDay() {
	day := times.Now2IntDay()
	fmt.Println(day > 0)

	// Output:
	// true
}

func ExampleParseIntDay() {
	day, err := times.ParseIntDay("20230918")

	fmt.Println(err)
	fmt.Println(day)
	fmt.Println(day.Year())
	fmt.Println(day.Month())
	fmt.Println(day.Day())

	// Output:
	// <nil>
	// 20230918
	// 2023
	// 9
	// 18
}
