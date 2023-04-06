package oss_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/oss"
)

// nolint: paralleltest
func TestAppendFile(t *testing.T) {
	ass := assert.New(t)

	patches1 := gomonkey.ApplyFuncReturn(os.OpenFile, nil, os.ErrClosed)
	defer patches1.Reset()

	_, err1 := oss.AppendFile(filepath.Join(os.TempDir(), "go-cli-test", "append.txt"))
	ass.NotNil(err1)

	patches2 := gomonkey.ApplyFuncReturn(os.MkdirAll, os.ErrClosed)
	defer patches2.Reset()

	patches3 := gomonkey.ApplyFuncReturn(os.IsNotExist, true)
	defer patches3.Reset()

	_, err2 := oss.AppendFile(filepath.Join(os.TempDir(), "go-cli-test", "append.txt"))
	ass.NotNil(err2)

	patches4 := gomonkey.ApplyFuncReturn(filepath.Abs, "", os.ErrClosed)
	defer patches4.Reset()

	_, err3 := oss.AppendFile(filepath.Join(os.TempDir(), "go-cli-test", "append.txt"))
	ass.NotNil(err3)
}
