package logs_test

import (
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
)

// nolint: paralleltest
func TestFileAbs(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(oss.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	ass := assert.New(t)
	_, err := logs.File(os.TempDir(), "test")

	ass.NotNil(err)
}

// nolint: paralleltest
func TestFileOpenFile(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(os.OpenFile, nil, os.ErrClosed)
	defer patches.Reset()

	ass := assert.New(t)
	_, err := logs.File(os.TempDir(), "test")

	ass.NotNil(err)
}

// nolint: paralleltest
func TestCloseFile(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(os.Remove, os.ErrClosed)
	defer patches.Reset()

	file := lo.Must1(os.CreateTemp(os.TempDir(), "test"))

	assert.NotNil(t, logs.CloseFile(file))
}
