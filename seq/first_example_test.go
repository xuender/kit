package seq_test

import (
	"fmt"

	"github.com/xuender/kit/v2/seq"
)

func ExampleFirst() {
	fmt.Println(seq.First(seq.Range(100)))
	fmt.Println(seq.First(seq.Range(0)))

	// Output:
	// 0 true
	// 0 false
}

func ExampleFirst2() {
	fmt.Println(seq.First2(seq.Range2(100)))
	fmt.Println(seq.First2(seq.Range2()))

	// Output:
	// 0 0 true
	// 0 0 false
}
