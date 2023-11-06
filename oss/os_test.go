package oss_test

import (
	"os/exec"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kit/oss"
)

// nolint: paralleltest
func TestOpen(t *testing.T) {
	req := require.New(t)
	cmd := &exec.Cmd{}

	patches1 := gomonkey.ApplyMethodReturn(cmd, "Start", nil)
	defer patches1.Reset()

	patches := gomonkey.ApplyFuncReturn(exec.Command, cmd)
	defer patches.Reset()

	req.NoError(oss.Open("file"))
	req.NoError(oss.Show("file"))

	pat := gomonkey.ApplyFuncReturn(oss.IsWindows, true)

	req.NoError(oss.Open("file"))
	req.NoError(oss.Show("."))
	req.Error(oss.Show("file"))
	req.NoError(oss.Show("os.go"))
	pat.Reset()

	pat = gomonkey.ApplyFuncReturn(oss.IsLinux, true)

	req.NoError(oss.Open("file"))
	req.NoError(oss.Show("file"))
	pat.Reset()

	pat = gomonkey.ApplyFuncReturn(oss.IsDarwin, true)

	req.NoError(oss.Open("file"))
	req.NoError(oss.Show("file"))
	pat.Reset()

	pat = gomonkey.ApplyFuncReturn(oss.IsLinux, false)

	req.Error(oss.Open("file"))
	req.Error(oss.Show("file"))
	pat.Reset()
}
