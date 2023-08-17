package cfg_test

import (
	"fmt"

	"github.com/xuender/kit/cfg"
)

func ExampleIsEncrypt() {
	fmt.Println(cfg.IsEncrypt("AES(A/43wTj2AVQboZZ0lNMqbw==)"))
	fmt.Println(cfg.IsEncrypt("xxx"))

	// Output:
	// true
	// false
}
