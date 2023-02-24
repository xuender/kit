package ios_test

import (
	"fmt"
	"io"

	"github.com/xuender/kit/ios"
)

func ExampleWrite() {
	fmt.Println(ios.Write(io.Discard, []byte("123"), []byte("4567")))

	// Output:
	// 7 <nil>
}
