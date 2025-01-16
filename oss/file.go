package oss

import (
	"io"
	"os"
	"path/filepath"

	"github.com/xuender/kit/v2/ops"
)

// AppendFile appends the given file.
func AppendFile(filename string) (*os.File, error) {
	path, err := Abs(filename)
	if err != nil {
		return nil, err
	}

	if dir := filepath.Dir(path); dir != "" && !Exist(dir) {
		if err := os.MkdirAll(dir, DefaultDirFileMod); err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultFileMode)
	if err != nil {
		return nil, err
	}

	return file, ops.Result2(file.Seek(0, io.SeekEnd))
}
