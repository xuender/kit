package oss

import (
	"os"
	"path/filepath"
)

// Remove 文件或目录，如果父目录为空则删除.
func Remove(path string, depth int) error {
	if !Exist(path) {
		return nil
	}

	if err := os.RemoveAll(path); err != nil {
		return err
	}

	for range depth {
		path = filepath.Dir(path)
		if dirs, err := os.ReadDir(path); err != nil || len(dirs) > 0 {
			return err
		}

		if err := os.Remove(path); err != nil {
			return err
		}
	}

	return nil
}
