package cfg

func Encrypt(str, key string) (string, error) {
	data, cipher, err := Parse(str)
	if err != nil {
		return "", err
	}

	return cipher.Stringify(cipher.EncryptBytes(data, key)), nil
}

func EncryptByCipher(src []byte, key string, cipher Cipher) string {
	return cipher.Stringify(cipher.EncryptBytes(src, key))
}

// Decrypt 解密.
func Decrypt(src, key string) (string, error) {
	srcBytes, cipher, err := Parse(src)
	if err != nil {
		return "", err
	}

	ret, err := cipher.DecryptBytes(srcBytes, key)
	if err != nil {
		return "", err
	}

	if _checkRegex.Match(ret) {
		return string(ret), err
	}

	return "", ErrKey
}
