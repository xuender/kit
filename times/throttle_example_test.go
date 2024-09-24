package times_test

import (
	"fmt"
	"time"

	"github.com/xuender/kit/times"
)

func ExampleThrottle() {
	play, wait := times.Throttle(func() { fmt.Println("play") }, time.Millisecond*100)
	defer wait()

	inter := times.Interval(time.Millisecond * 30)

	for range 10 {
		play()
		inter()
	}

	// Output:
	// play
	// play
	// play
}
