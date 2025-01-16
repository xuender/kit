package oss

import (
	"os"
	"path/filepath"
)

// Remove deletes the specified path and attempts to remove its parent directories up to a specified depth.
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
