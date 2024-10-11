package times_test

import (
	"context"
	"fmt"
	"time"

	"github.com/xuender/kit/times"
)

func ExampleDelayContext() {
	delay := times.DelayContext(time.Millisecond * 100)
	before := time.Now()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for num := range 3 {
		if delay(ctx) != nil {
			break
		}

		if num > 0 {
			cancel()
		}
	}

	fmt.Println(time.Since(before) >= time.Millisecond*200)

	// Output:
	// true
}
