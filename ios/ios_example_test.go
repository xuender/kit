package ios_test

import (
	"fmt"

	"github.com/xuender/kit/ios"
)

func ExampleWrite() {
	fmt.Println(ios.Write(ios.IgnoreWriter{}, []byte("123"), []byte("4567")))

	// Output:
	// 7 <nil>
}
