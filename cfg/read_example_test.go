package cfg_test

import (
	"fmt"

	"github.com/xuender/kit/cfg"
)

func ExampleRead() {
	data, err := cfg.Read([]byte(`a=AES(A/43wTj2AVQboZZ0lNMqbw==)
b=DES(LABOK5l6Q64=)
c=DES[abc]`), "key")

	fmt.Println(string(data))
	fmt.Println(err)

	// Output:
	// a=aaa
	// b=test2
	// c=abc
	// <nil>
}

func ExamplePathToBytes() {
	copyTest()

	_, err := cfg.PathToBytes("test.toml", "key")
	fmt.Println(err)

	// Output:
	// <nil>
}

func ExamplePathToReader() {
	copyTest()

	_, err := cfg.PathToReader("test.toml", "key")
	fmt.Println(err)

	// Output:
	// <nil>
}

func ExamplePathToString() {
	copyTest()

	_, err := cfg.PathToString("test.toml", "key")
	fmt.Println(err)

	// Output:
	// <nil>
}
