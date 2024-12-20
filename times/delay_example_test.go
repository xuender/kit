package times_test

import (
	"fmt"
	"time"

	"github.com/xuender/kit/times"
)

func ExampleDelay() {
	delay := times.Delay(time.Millisecond * 100)
	before := time.Now()

	for range 3 {
		delay()
	}

	fmt.Println(time.Since(before) >= time.Millisecond*300)

	// Output:
	// true
}
