package cfg_test

import (
	"fmt"

	"github.com/xuender/kit/cfg"
)

func ExampleEncrypt() {
	str, err := cfg.Encrypt("AES[123]", "password")
	fmt.Println(str[:4])
	fmt.Println(err)

	fmt.Println(cfg.Decrypt(str, "password"))

	// Output:
	// AES(
	// <nil>
	// 123 <nil>
}

func ExampleEncryptByCipher() {
	str := cfg.EncryptByCipher([]byte("123"), "password", cfg.AES)
	fmt.Println(str[:4])

	fmt.Println(cfg.Decrypt(str, "password"))
	// Output:
	// AES(
	// 123 <nil>
}

func ExampleDecrypt() {
	fmt.Println(cfg.Decrypt("AES(A/43wTj2AVQboZZ0lNMqbw==)", "key"))

	str := cfg.EncryptByCipher([]byte("123"), "", cfg.DES)
	fmt.Println(cfg.Decrypt(str, ""))
	fmt.Println(cfg.Decrypt(str, "err"))

	// Output:
	// aaa <nil>
	// 123 <nil>
	//  key error
}
