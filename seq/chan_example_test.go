package seq_test

import (
	"fmt"

	"github.com/xuender/kit/v2/seq"
)

// ExampleChan is an example function.
func ExampleChan() {
	cha := make(chan int, 3)
	cha <- 1
	cha <- 2
	cha <- 3

	close(cha)

	for num := range seq.Chan(cha) {
		if num > 1 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 1
}
