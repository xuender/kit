package times_test

import (
	"fmt"
	"time"

	"github.com/xuender/kit/times"
)

func ExampleDebounce() {
	play, wait := times.Debounce(func() { fmt.Println("play") }, time.Millisecond)
	defer wait()

	for range 10 {
		play()
	}

	// Output:
	// play
}
