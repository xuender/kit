package seq_test

import (
	"fmt"
	"time"

	"github.com/xuender/kit/v2/seq"
)

func ExampleCache() {
	input := func(yield func(int) bool) {
		for idx := range 5 {
			time.Sleep(time.Millisecond * 50)
			fmt.Println("input", idx)
			time.Sleep(time.Millisecond * 50)

			if !yield(idx) {
				return
			}
		}
	}

	for num := range seq.Cache(input, 10) {
		time.Sleep(time.Millisecond * 70)
		fmt.Println("output", num)
		time.Sleep(time.Millisecond * 70)

		if num > 3 {
			break
		}
	}

	// Output:
	// input 0
	// input 1
	// output 0
	// input 2
	// output 1
	// input 3
	// output 2
	// input 4
	// output 3
	// output 4
}
