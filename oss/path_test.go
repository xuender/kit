package oss_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/oss"
)

func TestAbs(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal("parent", filepath.Base(base.Must1(oss.Abs("parent/c/.."))))
	ass.NotEqual("~", filepath.Base(base.Must1(oss.Abs("~"))))

	home := base.Must1(os.UserHomeDir())
	path := base.Must1(oss.Abs("~"))
	ass.Equal(home, path)

	path = base.Must1(oss.Abs("~/file"))
	ass.True(strings.HasPrefix(path, home))

	path = base.Must1(oss.Abs("~file"))
	ass.NotEqual(len(home)+4, len(path))

	path = base.Must1(oss.Abs("~/../file"))
	ass.Equal(filepath.Join(filepath.Dir(home), "file"), path)
}

func TestAbs2(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	path := base.Must1(oss.Abs("./ff"))
	pwd := base.Must1(os.Getwd())
	ass.Equal(filepath.Join(pwd, "ff"), path)

	path = base.Must1(oss.Abs("../ff"))
	ass.Equal(filepath.Join(pwd, "..", "ff"), path)
}

// nolint
func TestAbs3(t *testing.T) {
	ass := assert.New(t)

	patches := gomonkey.ApplyFuncReturn(os.UserHomeDir, nil, os.ErrClosed)
	defer patches.Reset()

	ass.NotNil(oss.Abs("~"))
}

// nolint
func TestExist(t *testing.T) {
	ass := assert.New(t)

	ass.True(oss.Exist("doc.go"))
	ass.False(oss.Exist("unknown"))

	patches := gomonkey.ApplyFuncReturn(filepath.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	ass.False(oss.Exist(""))
}

// nolint
func TestIsDir(t *testing.T) {
	ass := assert.New(t)

	ass.False(oss.IsDir("doc.go"))
	ass.True(oss.IsDir("../oss"))

	patches := gomonkey.ApplyFuncReturn(filepath.Abs, nil, os.ErrClosed)
	defer patches.Reset()

	ass.False(oss.IsDir(""))
}

// nolint
func TestIsDir2(t *testing.T) {
	ass := assert.New(t)

	patches := gomonkey.ApplyFuncReturn(os.Stat, nil, os.ErrClosed)
	defer patches.Reset()

	ass.False(oss.IsDir("../oss"))
}
