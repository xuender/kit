package oss

import "runtime"

const (
	_windows = "windows"
	_linux   = "linux"
	_darwin  = "darwin"
)

// IsDarwin 是否是 Darwin 系统.
func IsDarwin() bool {
	return runtime.GOOS == _darwin
}

// IsLinux 是否是 Linux 系统.
func IsLinux() bool {
	return runtime.GOOS == _linux
}

// IsWindows 是否是 Windows 系统.
func IsWindows() bool {
	return runtime.GOOS == _windows
}
