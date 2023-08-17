package cfg

import (
	"bytes"
	"io"
	"os"

	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
)

// Read 秘文读取成明文.
func Read(data []byte, key string) ([]byte, error) {
	for _, src := range _regSecret.FindAll(data, -1) {
		value, err := Decrypt(string(src), key)
		if err != nil {
			return nil, err
		}

		data = bytes.ReplaceAll(data, src, []byte(value))
	}

	for _, src := range _regPlaintext.FindAll(data, -1) {
		value := src[4 : len(src)-1]
		data = bytes.ReplaceAll(data, src, value)
	}

	return data, nil
}

// Write 明文写入成秘文.
func Write(data []byte, key string) []byte {
	platintexts := _regPlaintext.FindAll(data, -1)
	if len(platintexts) == 0 {
		return nil
	}

	ret := make([]byte, len(data))
	copy(ret, data)

	for _, platintext := range platintexts {
		value, _ := Encrypt(string(platintext), key)

		ret = bytes.ReplaceAll(ret, platintext, []byte(value))
	}

	return ret
}

func Save(data []byte, path string) {
	if len(data) > 0 {
		logs.Log(os.WriteFile(path, data, oss.DefaultFileMode))
	}
}

func PathToBytes(path, key string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	Save(Write(data, key), path)

	return Read(data, key)
}

func PathToReader(path, key string) (io.Reader, error) {
	data, err := PathToBytes(path, key)
	if err != nil {
		return nil, err
	}

	Save(Write(data, key), path)

	return bytes.NewBuffer(data), nil
}

func PathToString(path, key string) (string, error) {
	data, err := PathToBytes(path, key)
	if err != nil {
		return "", err
	}

	Save(Write(data, key), path)

	return string(data), nil
}
