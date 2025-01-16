package times_test

import (
	"fmt"

	"github.com/xuender/kit/v2/ops"
	"github.com/xuender/kit/v2/times"
)

func ExampleNewParse() {
	parse := times.NewParse()

	fmt.Println(ops.Must(parse("2025-01-02")).Day())
	fmt.Println(ops.Must(parse("20210104")).Day())
	fmt.Println(ops.Must(parse("2024-01-01T16:14:13+08:00")).Minute())
	fmt.Println(ops.Must(parse("0103")).Day())
	fmt.Println(ops.Must(parse("0203")).Month())
	fmt.Println(ops.Must(parse("2025")).Year())

	_, err := parse("abcd")
	fmt.Println(err)

	// Output:
	// 2
	// 4
	// 14
	// 3
	// February
	// 2025
	// time format error
}
