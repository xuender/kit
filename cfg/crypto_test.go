package cfg_test

import (
	"testing"

	"github.com/xuender/kit/cfg"
)

func TestDecrypt(t *testing.T) {
	t.Parallel()

	if _, err := cfg.Decrypt("aaa(bb)", "key"); err == nil {
		t.Error("cipher error")
	}
}

func TestDecryptWith(t *testing.T) {
	t.Parallel()

	for i := 0; i < 10000; i++ {
		str := cfg.EncryptByCipher([]byte("阿弥陀佛"), "", cfg.DES)

		txt, err := cfg.Decrypt(str, "err")
		if err == nil {
			t.Error("error is nil", txt)
		}
	}
}

func TestPadding(t *testing.T) {
	t.Parallel()

	if data := cfg.Padding([]byte("1234"), 4); data[0] != '1' {
		t.Error("padding error")
	}
}

func TestEncrypt(t *testing.T) {
	t.Parallel()

	if _, err := cfg.Encrypt("aa[aa]", "key"); err == nil {
		t.Error("str err")
	}
}
