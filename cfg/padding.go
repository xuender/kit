package cfg

import (
	"bytes"
	"crypto/rand"
)

func pkcs5Trimming(encrypt []byte) ([]byte, error) {
	var (
		length    = len(encrypt)
		cipherLen = int(encrypt[length-1])
	)

	if length < cipherLen {
		return nil, ErrKey
	}

	return encrypt[:length-cipherLen], nil
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
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

func UnPadding(cipherText []byte) ([]byte, error) {
	var (
		length    = len(cipherText)
		cipherLen = int(cipherText[length-1])
	)

	if length < cipherLen {
		return nil, ErrKey
	}

	return cipherText[:length-cipherLen], nil
}

func getPaddingSize(cipherText []byte, blockSize int) int {
	remainder := len(cipherText) % blockSize
	if remainder == 0 {
		return blockSize
	}

	return blockSize - remainder
}
