// nolint: paralleltest
package cfg_test

import (
	"os"
	"testing"

	"github.com/xuender/kit/cfg"
	"github.com/xuender/kit/oss"
)

const (
	_file = "test.toml"
	_bak  = "test.bak"
)

func copyTest() {
	data, _ := os.ReadFile(_bak)
	_ = os.WriteFile(_file, data, oss.DefaultFileMode)
}

func TestPathToBytes(t *testing.T) {
	copyTest()

	if _, err := cfg.PathToBytes(_file, "key"); err != nil {
		t.Error(err)
	}

	copyTest()

	if _, err := cfg.PathToBytes(_file, "err"); err == nil {
		t.Error("key err")
	}

	copyTest()

	if _, err := cfg.PathToBytes("xxxx", "err"); err == nil {
		t.Error("file not found")
	}
}

func TestPathToReader(t *testing.T) {
	copyTest()

	if _, err := cfg.PathToReader(_file, "key"); err != nil {
		t.Error(err)
	}

	copyTest()

	if _, err := cfg.PathToReader(_file, "err"); err == nil {
		t.Error("key err")
	}
}

func TestPathToString(t *testing.T) {
	copyTest()

	if _, err := cfg.PathToString(_file, "key"); err != nil {
		t.Error(err)
	}

	copyTest()

	if _, err := cfg.PathToString(_file, "err"); err == nil {
		t.Error("key err")
	}
}
