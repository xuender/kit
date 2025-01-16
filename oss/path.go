package oss

import (
	"os"
	"path/filepath"
	"strings"
)

// Abs 返回路径的绝对表示，~符号表示用户根目录.
func Abs(path string) (string, error) {
	if path == "~" || strings.HasPrefix(path, string([]rune{'~', os.PathSeparator})) {
		home, err := os.UserHomeDir()
		if err != nil {
			return path, err
		}

		path = home + path[1:]
	}

	return filepath.Abs(path)
}

// Exist 文件或目录是否存在.
func Exist(path string) bool {
	path, err := Abs(path)
	if err != nil {
		return false
	}

	_, err = os.Stat(path)

	return !os.IsNotExist(err)
}

// IsDir 路径是否是目录.
func IsDir(path string) bool {
	path, err := Abs(path)
	if err != nil {
		return false
	}

	if info, err := os.Stat(path); err == nil {
		return info.IsDir()
	}

	return false
}
