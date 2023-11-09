package cfg

import (
	"crypto/aes"
	"crypto/cipher" // nolint
	"crypto/des"    // nolint
	"crypto/md5"    // nolint
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/xuender/kit/los"
)

type Cipher int

const (
	AES Cipher = iota
	DES
	AESMD5
	DESMD5
)

// nolint
var (
	ciphers = [...]Cipher{AES, DES}
	_names  = map[Cipher]string{AES: "AES", DES: "DES", AESMD5: "AESMD5", DESMD5: "DESMD5"}
)

func (p Cipher) Encrypt(src, key string) string {
	return base64.StdEncoding.EncodeToString(p.EncryptBytes([]byte(src), key))
}

func (p Cipher) EncryptBytes(src []byte, key string) []byte {
	blockMode, blockSize := p.Block(key, true)

	src = pkcs5Padding(src, blockSize)

	cryted := make([]byte, len(src))
	blockMode.CryptBlocks(cryted, src)

	return cryted
}

func (p Cipher) Decrypt(src, key string) (string, error) {
	data, err := p.DecryptBytes(los.Must(base64.StdEncoding.DecodeString(src)), key)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (p Cipher) DecryptBytes(src []byte, key string) ([]byte, error) {
	var (
		blockMode, _ = p.Block(key, false)
		orig         = make([]byte, len(src))
	)

	blockMode.CryptBlocks(orig, src)

	return pkcs5Trimming(orig)
}

func (p Cipher) String() string {
	return _names[p]
}

func (p Cipher) Block(key string, isEnc bool) (cipher.BlockMode, int) {
	var (
		keyBytes  []byte
		block     cipher.Block
		blockSize int
	)

	if p == AESMD5 || p == DESMD5 {
		tmp := md5.Sum([]byte(key)) // nolint
		keyBytes = tmp[:]
	} else {
		tmp := sha256.Sum256([]byte(key))
		keyBytes = tmp[:]
	}

	if p == DES || p == DESMD5 {
		blockSize = 8
		block = los.Must(des.NewCipher(keyBytes[:blockSize])) // nolint
	} else {
		block = los.Must(aes.NewCipher(keyBytes))
		blockSize = block.BlockSize()
	}

	if isEnc {
		return cipher.NewCBCEncrypter(block, keyBytes[:blockSize]), blockSize
	}

	return cipher.NewCBCDecrypter(block, keyBytes[:blockSize]), blockSize
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
