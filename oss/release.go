package oss

import (
	"os"
	"strings"
)

// IsRelease 是否是发布模式.
func IsRelease() bool {
	if os.Getenv("GIN_MODE") == "release" {
		return true
	}

	return !strings.HasPrefix(os.Args[0], os.TempDir())
}
