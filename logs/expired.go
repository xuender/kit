package logs

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/xuender/kit/oss"
)

// Expired 保留目录特定文件数量.
func Expired(path, name string, reserved int) error {
	path, err := oss.Abs(path)
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	// 历史日志
	var (
		ext  = filepath.Ext(name)
		olds = []string{}
		reg  = regexp.MustCompile(name[:len(name)-len(ext)] + `-\d{8}` + ext)
	)

	for _, entry := range entries {
		if reg.MatchString(entry.Name()) {
			olds = append(olds, entry.Name())
		}
	}

	num := len(olds) - reserved

	if num <= 0 {
		return nil
	}

	for _, delName := range olds[:num] {
		_ = os.Remove(filepath.Join(path, delName))
	}

	return nil
}
