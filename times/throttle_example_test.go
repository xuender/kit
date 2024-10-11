package times_test

import (
	"fmt"
	"time"

	"github.com/xuender/kit/times"
)

func ExampleThrottle() {
	play, wait := times.Throttle(func() { fmt.Println("play") }, time.Millisecond*100)
	defer wait()

	delay := times.Delay(time.Millisecond * 30)

	for range 10 {
		play()
		delay()
	}

	// Output:
	// play
	// play
	// play
}
