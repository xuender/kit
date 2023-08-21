package los_test

import (
	"fmt"

	"github.com/xuender/kit/los"
)

func ExampleSunday_IndexOf() {
	slice1 := []rune("abcaabcab")
	slice2 := []rune("abcab")
	sunday := los.NewSunday(slice2)
	fmt.Println(sunday.IndexOf(slice1))

	// Output:
	// 4
}
