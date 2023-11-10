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

func ExampleCipher_decrypt2() {
	fmt.Println(
		cfg.DES.Decrypt("HLZn0k/XGi5fmu2OCXsOuNdsstzF/2JCsY7q38hQEwSYFpGlO036ypruiYtMJmDB64JNZiiYh0sntLsphrX36A==", "woda"),
	)

	// Output:
	// TXprMU5tRTVZV1pqTW1Kak5ESTVNV0l4TkdNME4yRmhOMlppTURZM05HVQ== <nil>
}

func ExampleCipher_Decrypt() {
	fmt.Println(cfg.AESMD5.Decrypt(cfg.AESMD5.Encrypt("AESMD5", "pass"), "pass"))
	fmt.Println(cfg.DESMD5.Decrypt(cfg.DESMD5.Encrypt("DESMD5", "pass"), "pass"))
	fmt.Println(cfg.AES.Decrypt(cfg.AES.Encrypt("AES", "pass"), "pass"))
	fmt.Println(cfg.DES.Decrypt(cfg.DES.Encrypt("DES", "pass"), "pass"))

	// Output:
	// AESMD5 <nil>
	// DESMD5 <nil>
	// AES <nil>
	// DES <nil>
}
