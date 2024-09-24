package times_test

import (
	"fmt"
	"time"

	"github.com/xuender/kit/times"
)

func ExampleInterval() {
	inter := times.Interval(time.Millisecond * 100)
	before := time.Now()

	for range 3 {
		inter()
	}

	fmt.Println(time.Since(before) >= time.Millisecond*300)

	// Output:
	// true
}
