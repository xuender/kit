package logs_test

import (
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
)

type dirEntry struct{ name string }

func (p dirEntry) Name() string               { return p.name }
func (p dirEntry) IsDir() bool                { return false }
func (p dirEntry) Type() os.FileMode          { return oss.DefaultFileMode }
func (p dirEntry) Info() (os.FileInfo, error) { return nil, nil }

// nolint: paralleltest
func TestExpired(t *testing.T) {
	ass := assert.New(t)

	ass.Nil(logs.Expired(os.TempDir(), "test.log", 3))

	entries := []os.DirEntry{
		dirEntry{name: "test-23022611.log"},
		dirEntry{name: "test-23022612.log"},
		dirEntry{name: "test-23022613.log"},
		dirEntry{name: "test-23022614.log"},
		dirEntry{name: "test-23022616.log"},
		dirEntry{name: "test.log"},
	}

	patches := gomonkey.ApplyFuncReturn(os.ReadDir, entries, nil)
	defer patches.Reset()

	ass.Nil(logs.Expired(os.TempDir(), "test.log", 3))

	patches2 := gomonkey.ApplyFuncReturn(os.Remove, nil)
	defer patches2.Reset()

	ass.Nil(logs.Expired(os.TempDir(), "test.log", 3))
}

// nolint: paralleltest
func TestExpiredAbs(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(oss.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	ass := assert.New(t)

	ass.NotNil(logs.Expired(os.TempDir(), "test", 10))
}

// nolint: paralleltest
func TestExpiredReadDir(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(os.ReadDir, nil, os.ErrClosed)
	defer patches.Reset()

	ass := assert.New(t)

	ass.NotNil(logs.Expired(os.TempDir(), "test", 10))
}
