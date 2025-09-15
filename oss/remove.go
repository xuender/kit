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

	err := os.RemoveAll(path)
	if err != nil {
		return err
	}

	for range depth {
		path = filepath.Dir(path)

		dirs, err := os.ReadDir(path)
		if err != nil || len(dirs) > 0 {
			return err
		}

		err = os.Remove(path)
		if err != nil {
			return err
		}
	}

	return nil
}
