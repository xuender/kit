package logs_test

import (
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
)

// nolint: paralleltest
func TestSetLevel(t *testing.T) {
	ass := assert.New(t)

	logs.SetLevel(logs.Trace)
	ass.Equal(logs.Trace, logs.GetLevel())
	logs.SetLevel(logs.Debug)
	ass.Equal(logs.Debug, logs.GetLevel())
	logs.SetLevel(logs.Info)
	ass.Equal(logs.Info, logs.GetLevel())
	logs.SetLevel(logs.Warn)
	ass.Equal(logs.Warn, logs.GetLevel())
	logs.SetLevel(logs.Error)
	ass.Equal(logs.Error, logs.GetLevel())
	logs.D.Println("pass")
	logs.E.Println("pass")
}

// nolint: paralleltest
func TestSetLogFile(_ *testing.T) {
	_ = logs.SetLogFile(os.TempDir(), "test")
	defer logs.Close()

	_ = logs.SetLogFile(os.TempDir(), "test1")
}

// nolint: paralleltest
func TestSetLogFileError(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(oss.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	req := require.New(t)

	req.Error(logs.SetLogFile(os.TempDir(), "test"))
}

// nolint: paralleltest
func TestSetTrace(_ *testing.T) {
	_ = logs.SetTraceFile(os.TempDir(), "test")
	defer logs.Close()

	_ = logs.SetTraceFile(os.TempDir(), "test1")
}

// nolint: paralleltest
func TestSetTraceFile(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(oss.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	req := require.New(t)
	req.Error(logs.SetTraceFile(os.TempDir(), "test"))
}

// nolint: paralleltest
func TestSetDebug(_ *testing.T) {
	_ = logs.SetDebugFile(os.TempDir(), "test")
	defer logs.Close()

	_ = logs.SetDebugFile(os.TempDir(), "test1")
}

// nolint: paralleltest
func TestSetDebugFile(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(oss.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	req := require.New(t)
	req.Error(logs.SetDebugFile(os.TempDir(), "test"))
}

// nolint: paralleltest
func TestSetInfo(_ *testing.T) {
	_ = logs.SetInfoFile(os.TempDir(), "test")
	defer logs.Close()

	_ = logs.SetInfoFile(os.TempDir(), "test1")
}

// nolint: paralleltest
func TestSetInfoFile(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(oss.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	req := require.New(t)
	req.Error(logs.SetInfoFile(os.TempDir(), "test"))
}

// nolint: paralleltest
func TestSetWarn(_ *testing.T) {
	_ = logs.SetWarnFile(os.TempDir(), "test")
	defer logs.Close()

	_ = logs.SetWarnFile(os.TempDir(), "test1")
}

// nolint: paralleltest
func TestSetWarnFile(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(oss.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	req := require.New(t)
	req.Error(logs.SetWarnFile(os.TempDir(), "test"))
}

// nolint: paralleltest
func TestSetError(_ *testing.T) {
	_ = logs.SetErrorFile(os.TempDir(), "test")
	defer logs.Close()

	_ = logs.SetErrorFile(os.TempDir(), "test1")
	logs.E.Println("error")
}

// nolint: paralleltest
func TestSetErrorFile(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(oss.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	req := require.New(t)
	req.Error(logs.SetErrorFile(os.TempDir(), "test"))
}

// nolint: paralleltest
func TestRotating(_ *testing.T) {
	logs.Rotating(os.TempDir(), "test.log", func(_, _ string) error {
		return os.ErrClosed
	})()

	old := logs.RetentionDays
	logs.RetentionDays = 0

	logs.Rotating(os.TempDir(), "test.log", func(_, _ string) error {
		return nil
	})()

	logs.RetentionDays = old

	patches := gomonkey.ApplyFuncReturn(oss.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	logs.Rotating(os.TempDir(), "test.log", func(_, _ string) error {
		return nil
	})()
}
