package cfg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des" // nolint
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

type Cipher int

const (
	AES Cipher = iota
	DES
)

// nolint
var ciphers = [...]Cipher{AES, DES}

func (p Cipher) String() string {
	if p == DES {
		return "DES"
	}

	return "AES"
}

func (p Cipher) Block(key string) cipher.Block {
	keyBytes := sha256.Sum256([]byte(key))
	if p == DES {
		// nolint
		block, _ := des.NewCipher(keyBytes[:8])

		return block
	}

	block, _ := aes.NewCipher(keyBytes[:])

	return block
}

func (p Cipher) Stringify(data []byte) string {
	return fmt.Sprintf("%s(%s)", p.String(), base64.StdEncoding.EncodeToString(data))
}

// IsEncrypt 是否加密.
func IsEncrypt(str string) bool {
	_, _, err := Parse(str)

	return err == nil
}

// Parse 解析密文, 返回数据和加密算法.
func Parse(str string) ([]byte, Cipher, error) {
	var (
		cipher      Cipher
		isSecret    = _secretRegex.MatchString(str)
		isPlaintext = _plaintextRegex.MatchString(str)
	)

	if !isSecret && !isPlaintext {
		return nil, cipher, ErrNoCipher
	}

	for _, cip := range ciphers {
		if strings.HasPrefix(str, cip.String()) {
			cipher = cip
		}
	}

	text := str[4 : len(str)-1]
	if isSecret {
		src, err := base64.StdEncoding.DecodeString(text)

		return src, cipher, err
	}

	return []byte(text), cipher, nil
}
