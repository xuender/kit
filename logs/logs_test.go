package logs_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/logs"
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
func TestSetLogFile(t *testing.T) {
	_ = logs.SetLogFile(os.TempDir(), "test")
	defer logs.Close()

	_ = logs.SetLogFile(os.TempDir(), "test1")
}

// nolint: paralleltest
func TestSetTrace(t *testing.T) {
	_ = logs.SetTraceFile(os.TempDir(), "test")
	defer logs.Close()

	_ = logs.SetTraceFile(os.TempDir(), "test1")
}

// nolint: paralleltest
func TestSetDebug(t *testing.T) {
	_ = logs.SetDebugFile(os.TempDir(), "test")
	defer logs.Close()

	_ = logs.SetDebugFile(os.TempDir(), "test1")
}

// nolint: paralleltest
func TestSetInfo(t *testing.T) {
	_ = logs.SetInfoFile(os.TempDir(), "test")
	defer logs.Close()

	_ = logs.SetInfoFile(os.TempDir(), "test1")
}

// nolint: paralleltest
func TestSetWarn(t *testing.T) {
	_ = logs.SetWarnFile(os.TempDir(), "test")
	defer logs.Close()

	_ = logs.SetWarnFile(os.TempDir(), "test1")
}

// nolint: paralleltest
func TestSetError(t *testing.T) {
	_ = logs.SetErrorFile(os.TempDir(), "test")
	defer logs.Close()

	_ = logs.SetErrorFile(os.TempDir(), "test1")
	logs.E.Println("error")
}
