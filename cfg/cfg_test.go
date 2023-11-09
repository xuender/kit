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
	_file = "test.toml"
	_data = `k1 = "AES(oY8aex0d4WjWokMZSUDyDQ==)"
k2 = "DES(Jiga4xHtvWM=)"
k3 = "DES[中文]"`
	_keyErr = "key err"
)

func TestCfg_Bytes(t *testing.T) {
	patchRead := gomonkey.ApplyFuncReturn(os.ReadFile, []byte(_data), nil)
	defer patchRead.Reset()

	patchWrite := gomonkey.ApplyFuncReturn(os.WriteFile, nil)
	defer patchWrite.Reset()

	if data, err := cfg.New("key").Bytes(_file); err == nil {
		if !bytes.Contains(data, []byte("aaa")) {
			t.Error("miss aaa")
		}

		if !bytes.Contains(data, []byte("中文")) {
			t.Error("miss 中文")
		}
	} else {
		t.Error(err)
	}

	if _, err := cfg.New("err").Bytes(_file); err == nil {
		t.Error(_keyErr)
	}
}

func TestCfg_BytesError(t *testing.T) {
	patchRead := gomonkey.ApplyFuncReturn(os.ReadFile, nil, cfg.ErrKey)
	defer patchRead.Reset()

	patchWrite := gomonkey.ApplyFuncReturn(os.WriteFile, nil)
	defer patchWrite.Reset()

	if _, err := cfg.New("err").Bytes(_file); err == nil {
		t.Error("file not found")
	}
}

func TestCfg_Reader(t *testing.T) {
	patchRead := gomonkey.ApplyFuncReturn(os.ReadFile, []byte(_data), nil)
	defer patchRead.Reset()

	patchWrite := gomonkey.ApplyFuncReturn(os.WriteFile, nil)
	defer patchWrite.Reset()

	if _, err := cfg.New("key").Reader(_file); err != nil {
		t.Error(err)
	}

	if _, err := cfg.New("err").Reader(_file); err == nil {
		t.Error(_keyErr)
	}
}

func TestPathToString(t *testing.T) {
	patchRead := gomonkey.ApplyFuncReturn(os.ReadFile, []byte(_data), nil)
	defer patchRead.Reset()

	patchWrite := gomonkey.ApplyFuncReturn(os.WriteFile, nil)
	defer patchWrite.Reset()

	if _, err := cfg.New("key").String(_file); err != nil {
		t.Error(err)
	}

	if _, err := cfg.New("err").String(_file); err == nil {
		t.Error(_keyErr)
	}
}
