package seq_test

import (
	"fmt"
	"time"

	"github.com/xuender/kit/v2/seq"
)

func ExampleSplitGroup() {
	iter, done := seq.SplitGroup([]int{1, 3, 5, 7, 9, 2, 4, 6, 8}, func(num int) int { return num % 2 })

	for num := range iter {
		fmt.Println(num)

		go func(val int) {
			time.Sleep(time.Millisecond * time.Duration(val*10))
			done(val)
		}(num)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
}
