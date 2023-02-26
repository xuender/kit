package logs

import (
	"os"
	"strconv"
	"strings"
)

// Level 日志级别.
type Level int

const (
	Trace Level = iota
	Debug
	Info
	Warn
	Error
)

// nolint: gochecknoglobals
var _levels = [...]string{"Trace", "Debug", "Info", "Warn", "Error"}

// DefaultLevel 默认日志级别.
func DefaultLevel() Level {
	level := os.Getenv("LOGS_LEVEL")

	for index, label := range _levels {
		if strings.EqualFold(level, label) || level == strconv.Itoa(index) {
			return Level(index)
		}
	}

	return Debug
}
