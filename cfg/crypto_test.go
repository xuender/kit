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

	for range 10000 {
		str := cfg.EncryptByCipher([]byte("阿弥陀佛"), "", cfg.DES)

		txt, err := cfg.Decrypt(str, "err")
		if err == nil {
			t.Error("error is nil", txt)
		}
	}
}

func TestEncrypt(t *testing.T) {
	t.Parallel()

	if _, err := cfg.Encrypt("aa[aa]", "key"); err == nil {
		t.Error("str err")
	}
}
