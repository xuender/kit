package los_test

import (
	"fmt"

	"github.com/xuender/kit/los"
)

func ExampleSunday_IndexOf() {
	sub := []rune("abcab")
	sunday := los.NewSunday(sub)
	fmt.Println(sunday.IndexOf([]rune("abcaabcab")))

	// Output:
	// 4
}
