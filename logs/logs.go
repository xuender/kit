package logs

import (
	"io"
	"os"

	"github.com/xuender/kit/base"
	"github.com/xuender/kit/times"
)

// nolint: gochecknoglobals, varnamelen
var (
	_loggers = [...]*logger{
		{output: _colorTrace},
		{output: _colorDebug},
		{output: os.Stderr},
		{output: _colorWarn},
		{output: _colorError},
	}
	_level Level
	// T 跟踪.
	T = _loggers[Trace].newLog("[T] ")
	// D 调试.
	D = _loggers[Debug].newLog("[D] ")
	// I 消息.
	I = _loggers[Info].newLog("[I] ")
	// W 警告.
	W = _loggers[Warn].newLog("[W] ")
	// E 错误.
	E = _loggers[Error].newLog("[E] ")
	// RetentionDays 日志保留天数，0为永久.
	RetentionDays = base.Seven
)

// nolint: gochecknoinits
func init() {
	SetLevel(DefaultLevel())
}

// SetLog 默认输出.
func SetLog(writer io.Writer) {
	for _, logger := range _loggers {
		logger.setOutput(writer)
	}

	SetLevel(_level)
}

// SetLogFile 默认文件输出.
func SetLogFile(path, file string) error {
	writer, err := File(path, file)
	if err != nil {
		return err
	}

	SetLog(writer)
	rotate(path, file, SetLogFile)

	return nil
}

func rotate(path, file string, yield func(string, string) error) {
	times.Hour(Rotating(path, file, yield))
}

// Rotating 返回旧日志删除方法.
func Rotating(path, file string, yield func(string, string) error) func() {
	return func() {
		if err := yield(path, file); err != nil {
			E.Println(err)

			return
		}

		if RetentionDays <= 0 {
			return
		}

		Log(Expired(path, file, RetentionDays*base.TwentyFour))
	}
}

// SetTrace 设置跟踪.
func SetTrace(writer io.Writer) {
	_loggers[Trace].setOutput(writer)

	SetLevel(_level)
}

// SetTraceFile 设置跟踪文件.
func SetTraceFile(path, file string) error {
	writer, err := File(path, file)
	if err != nil {
		return err
	}

	SetTrace(writer)
	rotate(path, file, SetTraceFile)

	return nil
}

// SetDebug 设置调试.
func SetDebug(writer io.Writer) {
	_loggers[Debug].setOutput(writer)

	SetLevel(_level)
}

// SetDebugFile 设置调试输出.
func SetDebugFile(path, file string) error {
	writer, err := File(path, file)
	if err != nil {
		return err
	}

	SetDebug(writer)
	rotate(path, file, SetDebugFile)

	return nil
}

// SetInfo 设置信息.
func SetInfo(writer io.Writer) {
	_loggers[Info].setOutput(writer)

	SetLevel(_level)
}

// SetInfoFile 设置文件输出.
func SetInfoFile(path, file string) error {
	writer, err := File(path, file)
	if err != nil {
		return err
	}

	SetInfo(writer)
	rotate(path, file, SetInfoFile)

	return nil
}

// SetWarn 设置警告.
func SetWarn(writer io.Writer) {
	_loggers[Warn].setOutput(writer)

	SetLevel(_level)
}

// SetWarnFile 设置警告文件输出.
func SetWarnFile(path, file string) error {
	writer, err := File(path, file)
	if err != nil {
		return err
	}

	SetWarn(writer)
	rotate(path, file, SetWarnFile)

	return nil
}

// SetError 设置错误.
func SetError(writer io.Writer) {
	_loggers[Error].setOutput(writer)

	SetLevel(_level)
}

// SetErrorFile 设置错误文件输出.
func SetErrorFile(path, file string) error {
	writer, err := File(path, file)
	if err != nil {
		return err
	}

	SetError(writer)
	rotate(path, file, SetErrorFile)

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

// Log 输出日志.
func Log(values ...any) {
	if base.AllNil(values...) {
		return
	}
	// 包含error 则使用Error输出
	for _, value := range values {
		if _, ok := value.(error); ok {
			E.Println(values...)

			return
		}
	}

	I.Println(values...)
}
