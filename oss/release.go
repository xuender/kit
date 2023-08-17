package oss

import (
	"os"
	"path/filepath"
	"strings"
)

// IsRelease 是否是发布模式.
func IsRelease() bool {
	if Version != "" {
		return true
	}

	if base := filepath.Base(os.Args[0]); strings.HasPrefix(base, "__") {
		return false
	}

	return !strings.HasPrefix(os.Args[0], os.TempDir())
}
