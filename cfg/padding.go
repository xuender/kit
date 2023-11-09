package cfg

import "bytes"

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
