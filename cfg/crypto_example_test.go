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

func ExampleDecrypt_md5() {
	fmt.Println(cfg.AESMD5.Decrypt("lob52vO/Av/yk0Ty+DBDag==", "pass"))

	// Output:
	// abc <nil>
}
