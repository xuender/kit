// nolint: paralleltest
package cfg_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/xuender/kit/cfg"
)

const (
	_data = `k1 = "AES(A/43wTj2AVQboZZ0lNMqbw==)"
k2 = "DES(LABOK5l6Q64=)"
k3 = "DES[中文]"`
	_file   = "test.toml"
	_keyErr = "key err"
)

func TestCfg_Bytes(t *testing.T) {
	patchRead := gomonkey.ApplyFuncReturn(os.ReadFile, []byte(_data), nil)
	defer patchRead.Reset()

	patchWrite := gomonkey.ApplyFuncReturn(os.WriteFile, nil)
	defer patchWrite.Reset()

	data, err := cfg.New("key").Bytes(_file)
	if err == nil {
		if !bytes.Contains(data, []byte("aaa")) {
			t.Error("miss aaa")
		}

		if !bytes.Contains(data, []byte("中文")) {
			t.Error("miss 中文")
		}
	} else {
		t.Error(err)
	}

	_, err = cfg.New("err").Bytes(_file)
	if err == nil {
		t.Error(_keyErr)
	}
}

func TestCfg_BytesError(t *testing.T) {
	patchRead := gomonkey.ApplyFuncReturn(os.ReadFile, nil, cfg.ErrKey)
	defer patchRead.Reset()

	patchWrite := gomonkey.ApplyFuncReturn(os.WriteFile, nil)
	defer patchWrite.Reset()

	_, err := cfg.New("err").Bytes(_file)
	if err == nil {
		t.Error("file not found")
	}
}

func TestCfg_Reader(t *testing.T) {
	patchRead := gomonkey.ApplyFuncReturn(os.ReadFile, []byte(_data), nil)
	defer patchRead.Reset()

	patchWrite := gomonkey.ApplyFuncReturn(os.WriteFile, nil)
	defer patchWrite.Reset()

	_, err := cfg.New("key").Reader(_file)
	if err != nil {
		t.Error(err)
	}

	_, err = cfg.New("err").Reader(_file)
	if err == nil {
		t.Error(_keyErr)
	}
}

func TestPathToString(t *testing.T) {
	patchRead := gomonkey.ApplyFuncReturn(os.ReadFile, []byte(_data), nil)
	defer patchRead.Reset()

	patchWrite := gomonkey.ApplyFuncReturn(os.WriteFile, nil)
	defer patchWrite.Reset()

	_, err := cfg.New("key").String(_file)
	if err != nil {
		t.Error(err)
	}

	_, err = cfg.New("err").String(_file)
	if err == nil {
		t.Error(_keyErr)
	}
}
