package times_test

import (
	"fmt"
	"time"

	"github.com/xuender/kit/times"
)

func ExampleDebounce() {
	play, wait := times.Debounce(func() { fmt.Println("play") }, time.Millisecond)

	for range 10 {
		play()
	}

	wait()

	// Output:
	// play
}
