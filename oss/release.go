package oss

import (
	"os"
	"strings"
)

// IsRelease 是否是发布模式.
func IsRelease() bool {
	if Version != "" {
		return true
	}

	return !strings.HasPrefix(os.Args[0], os.TempDir())
}
