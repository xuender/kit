package hash_test

import (
	"fmt"

	"github.com/xuender/kit/hash"
)

func ExampleSipHash128() {
	fmt.Println(hash.SipHash128([]byte("123")))

	// Output:
	// 8693645449139915215 11618447955228391416
}

func ExampleSipHash64() {
	fmt.Println(hash.SipHash64([]byte("123")))

	// Output:
	// 9379172312344772015
}

func ExampleSipHashNumber() {
	fmt.Println(hash.SipHashNumber([]byte("123")))
	// output:
	// 2677888159399343
}

func ExampleSipHashHex() {
	fmt.Println(hash.SipHashHex([]byte("123")))
	// output:
	// 822983866c7d3daf
}
