package seq_test

import (
	"fmt"

	"github.com/xuender/kit/v2/seq"
)

func ExampleLast() {
	fmt.Println(seq.Last(seq.Range(10)))

	// Output:
	// 9 true
}

func ExampleLast2() {
	fmt.Println(seq.Last2(seq.Range2(10)))

	// Output:
	// 9 9 true
}
