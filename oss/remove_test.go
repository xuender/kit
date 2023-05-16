package oss_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/oss"
	"github.com/xuender/kit/times"
)

func TestRemove(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	worker := times.NewIDWorkerByKey("a")
	dir := filepath.Join(os.TempDir(), worker.String())
	path := filepath.Join(dir, worker.String(), worker.String(), worker.String())

	ass.Nil(os.MkdirAll(path, oss.DefaultDirFileMod))
	ass.Nil(os.WriteFile(filepath.Join(path, worker.String()), []byte("test"), oss.DefaultFileMode))
	ass.Nil(os.WriteFile(filepath.Join(path, worker.String()), []byte("test"), oss.DefaultFileMode))
	ass.Nil(os.WriteFile(filepath.Join(path, worker.String()), []byte("test"), oss.DefaultFileMode))
	ass.Nil(oss.Remove(path, 3))
	ass.False(oss.Exist(dir))
	ass.Nil(oss.Remove(filepath.Join(os.TempDir(), worker.String()), 10))
}

func TestRemove2(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	worker := times.NewIDWorkerByKey("b")
	dir := filepath.Join(os.TempDir(), worker.String())
	path := filepath.Join(dir, worker.String(), worker.String(), worker.String())

	ass.Nil(os.MkdirAll(path, oss.DefaultDirFileMod))
	ass.Nil(os.WriteFile(filepath.Join(path, worker.String()), []byte("test"), oss.DefaultFileMode))
	ass.Nil(os.WriteFile(filepath.Join(filepath.Dir(path), worker.String()), []byte("test"), oss.DefaultFileMode))

	ass.Nil(oss.Remove(path, 3))
	ass.True(oss.Exist(dir))
}

// nolint: paralleltest
func TestRemove3(t *testing.T) {
	ass := assert.New(t)
	worker := times.NewIDWorkerByKey("c")
	dir := filepath.Join(os.TempDir(), worker.String())
	path := filepath.Join(dir, worker.String(), worker.String(), worker.String())

	ass.Nil(os.MkdirAll(path, oss.DefaultDirFileMod))
	file := filepath.Join(path, worker.String())
	ass.Nil(os.WriteFile(file, []byte("test"), oss.DefaultFileMode))

	patches1 := gomonkey.ApplyFuncReturn(os.Remove, assert.AnError)
	defer patches1.Reset()

	ass.NotNil(oss.Remove(path, 3))

	ass.Nil(os.MkdirAll(path, oss.DefaultDirFileMod))
	ass.Nil(os.WriteFile(file, []byte("test"), oss.DefaultFileMode))

	patches2 := gomonkey.ApplyFuncReturn(os.RemoveAll, assert.AnError)
	defer patches2.Reset()

	ass.NotNil(oss.Remove(path, 3))
}
