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

func (p Cipher) EncodeToString(src, key string) string {
	return fmt.Sprintf("%s(%s)", p.String(), p.Encrypt(src, key))
}

func (p Cipher) EncryptBytes(src []byte, key string) []byte {
	if p == AESMD5 || p == DESMD5 {
		blockMode, blockSize := p.BlockMode(key, true)

		src = pkcs5Padding(src, blockSize)

		cryted := make([]byte, len(src))
		blockMode.CryptBlocks(cryted, src)

		return cryted
	}

	var (
		ret      []byte
		block    = p.Block(key)
		srcBytes = Padding(src, block.BlockSize())
		tmp      = make([]byte, block.BlockSize())
	)

	for index := 0; index < len(srcBytes); index += block.BlockSize() {
		block.Encrypt(tmp, srcBytes[index:index+block.BlockSize()])
		ret = append(ret, tmp...)
	}

	return ret
}

func (p Cipher) Decrypt(src, key string) (string, error) {
	data, err := p.DecryptBytes(los.Must(base64.StdEncoding.DecodeString(src)), key)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (p Cipher) DecryptBytes(src []byte, key string) ([]byte, error) {
	if p == AESMD5 || p == DESMD5 {
		var (
			blockMode, _ = p.BlockMode(key, false)
			orig         = make([]byte, len(src))
		)

		blockMode.CryptBlocks(orig, src)

		return pkcs5Trimming(orig)
	}

	var (
		ret   []byte
		block = p.Block(key)
		tmp   = make([]byte, block.BlockSize())
	)

	for index := 0; index < len(src); index += block.BlockSize() {
		block.Decrypt(tmp, src[index:index+block.BlockSize()])
		ret = append(ret, tmp...)
	}

	ret, err := UnPadding(ret)
	if err != nil {
		return nil, err
	}

	if _checkRegex.Match(ret) {
		return ret, nil
	}

	return nil, ErrKey
}

func (p Cipher) String() string {
	return _names[p]
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

func (p Cipher) BlockMode(key string, isEnc bool) (cipher.BlockMode, int) {
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
