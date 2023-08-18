package cfg

import (
	"bytes"
	"io"
	"os"

	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
)

type Cfg struct {
	key string
}

func New(password string) *Cfg {
	return &Cfg{key: password}
}

// Bytes 配置转字节.
func (p *Cfg) Bytes(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	save(p.Write(data), path)

	return p.Read(data)
}

// Reader 配置转 io.Reader .
func (p *Cfg) Reader(path string) (io.Reader, error) {
	data, err := p.Bytes(path)
	if err != nil {
		return nil, err
	}

	save(p.Write(data), path)

	return bytes.NewBuffer(data), nil
}

// String 配置转字符串.
func (p *Cfg) String(path string) (string, error) {
	data, err := p.Bytes(path)
	if err != nil {
		return "", err
	}

	save(p.Write(data), path)

	return string(data), nil
}

// Read 秘文读取成明文.
func (p *Cfg) Read(data []byte) ([]byte, error) {
	for _, src := range _regSecret.FindAll(data, -1) {
		value, err := Decrypt(string(src), p.key)
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
func (p *Cfg) Write(data []byte) []byte {
	platintexts := _regPlaintext.FindAll(data, -1)
	if len(platintexts) == 0 {
		return nil
	}

	ret := make([]byte, len(data))
	copy(ret, data)

	for _, platintext := range platintexts {
		value, _ := Encrypt(string(platintext), p.key)

		ret = bytes.ReplaceAll(ret, platintext, []byte(value))
	}

	return ret
}

func save(data []byte, path string) {
	if len(data) > 0 {
		logs.Log(os.WriteFile(path, data, oss.DefaultFileMode))
	}
}
