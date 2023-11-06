package logs_test

import (
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
)

// nolint: paralleltest
func TestFileAbs(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(oss.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	req := require.New(t)
	_, err := logs.File(os.TempDir(), "test")

	req.Error(err)
}

// nolint: paralleltest
func TestFileOpenFile(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(os.OpenFile, nil, os.ErrClosed)
	defer patches.Reset()

	req := require.New(t)
	_, err := logs.File(os.TempDir(), "test")

	req.Error(err)
}

// nolint: paralleltest
func TestCloseFile(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(os.Remove, os.ErrClosed)
	defer patches.Reset()

	req := require.New(t)
	file := lo.Must1(os.CreateTemp(os.TempDir(), "test"))

	req.Error(logs.CloseFile(file))
}
