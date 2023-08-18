package cfg

import (
	"crypto/rand"
)

func Encrypt(str, key string) (string, error) {
	data, cipher, err := Parse(str)
	if err != nil {
		return "", err
	}

	return EncryptByCipher(data, key, cipher), nil
}

// EncryptByCipher 加密.
func EncryptByCipher(str []byte, key string, cipher Cipher) string {
	var (
		ret      []byte
		block    = cipher.Block(key)
		srcBytes = Padding(str, block.BlockSize())
		tmp      = make([]byte, block.BlockSize())
	)

	for index := 0; index < len(srcBytes); index += block.BlockSize() {
		block.Encrypt(tmp, srcBytes[index:index+block.BlockSize()])
		ret = append(ret, tmp...)
	}

	return cipher.Stringify(ret)
}

// Decrypt 解密.
func Decrypt(src, key string) (string, error) {
	var ret []byte

	srcBytes, cipher, err := Parse(src)
	if err != nil {
		return "", err
	}

	var (
		block = cipher.Block(key)
		tmp   = make([]byte, block.BlockSize())
	)

	for index := 0; index < len(srcBytes); index += block.BlockSize() {
		block.Decrypt(tmp, srcBytes[index:index+block.BlockSize()])
		ret = append(ret, tmp...)
	}

	text, err := UnPadding(ret)
	if err != nil {
		return "", err
	}

	if _regCheck.MatchString(text) {
		return text, err
	}

	return "", ErrKey
}

func Padding(cipherText []byte, blockSize int) []byte {
	var (
		padding = getPaddingSize(cipherText, blockSize)
		padData = make([]byte, padding-1)
	)

	_, _ = rand.Read(padData)
	// nolint
	padData = append(padData, byte(padding))

	return append(cipherText, padData...)
}

func UnPadding(cipherText []byte) (string, error) {
	var (
		length    = len(cipherText)
		cipherLen = int(cipherText[length-1])
	)

	if length < cipherLen {
		return "", ErrKey
	}

	return string(cipherText[:length-cipherLen]), nil
}

func getPaddingSize(cipherText []byte, blockSize int) int {
	remainder := len(cipherText) % blockSize
	if remainder == 0 {
		return blockSize
	}

	return blockSize - remainder
}
