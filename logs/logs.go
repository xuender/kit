package logs

import (
	"io"
	"os"
)

// nolint: gochecknoglobals
var (
	_loggers = [...]*logger{
		{output: _colorTrace},
		{output: _colorDebug},
		{output: os.Stderr},
		{output: _colorWarn},
		{output: _colorError},
	}
	_level = Info
	// T 跟踪.
	T = _loggers[Trace].newLog("[T] ", true)
	// D 调试.
	D = _loggers[Debug].newLog("[D] ", true)
	// I 消息.
	I = _loggers[Info].newLog("[I] ", false)
	// W 警告.
	W = _loggers[Warn].newLog("[W] ", false)
	// E 错误.
	E = _loggers[Error].newLog("[E] ", false)
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

// SetLog 默认输出.
func SetLog(writer io.Writer) {
	for _, logger := range _loggers {
		logger.setOutput(writer)
	}

	SetLevel(_level)
}

// SetLogFile 默认文件输出.
func SetLogFile(path, file string) error {
	writer, err := LogFile(path, file)
	if err != nil {
		return err
	}

	SetLog(writer)

	return nil
}

// SetTrace 设置跟踪.
func SetTrace(writer io.Writer) {
	_loggers[Trace].setOutput(writer)

	SetLevel(_level)
}

// SetTraceFile 设置跟踪文件.
func SetTraceFile(path, file string) error {
	writer, err := LogFile(path, file)
	if err != nil {
		return err
	}

	SetTrace(writer)

	return nil
}

// SetDebug 设置调试.
func SetDebug(writer io.Writer) {
	_loggers[Debug].setOutput(writer)

	SetLevel(_level)
}

// SetDebugFile 设置调试输出.
func SetDebugFile(path, file string) error {
	writer, err := LogFile(path, file)
	if err != nil {
		return err
	}

	SetDebug(writer)

	return nil
}

// SetInfo 设置信息.
func SetInfo(writer io.Writer) {
	_loggers[Info].setOutput(writer)

	SetLevel(_level)
}

// SetInfoFile 设置文件输出.
func SetInfoFile(path, file string) error {
	writer, err := LogFile(path, file)
	if err != nil {
		return err
	}

	SetInfo(writer)

	return nil
}

// SetWarn 设置警告.
func SetWarn(writer io.Writer) {
	_loggers[Warn].setOutput(writer)

	SetLevel(_level)
}

// SetWarnFile 设置警告文件输出.
func SetWarnFile(path, file string) error {
	writer, err := LogFile(path, file)
	if err != nil {
		return err
	}

	SetWarn(writer)

	return nil
}

// SetError 设置错误.
func SetError(writer io.Writer) {
	_loggers[Error].setOutput(writer)

	SetLevel(_level)
}

// SetErrorFile 设置错误文件输出.
func SetErrorFile(path, file string) error {
	writer, err := LogFile(path, file)
	if err != nil {
		return err
	}

	SetError(writer)

	return nil
}

// SetLevel 设置日志级别.
func SetLevel(level Level) {
	for _, logger := range _loggers {
		logger.reset()
	}

	for i := Trace; i < level; i++ {
		_loggers[i].ignore()
	}

	_level = level
}

// GetLevel 获取日志级别.
func GetLevel() Level {
	return _level
}
