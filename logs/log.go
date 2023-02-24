package logs

import (
	"io"
	"log"
	"os"
)

// nolint: gochecknoglobals, varnamelen
var (
	_passer  = &passer{}
	_writers = [...]io.Writer{_colorTrace, _colorDebug, os.Stderr, _colorWarn, _colorError}
	_workers = [...]io.Writer{_passer, _passer, os.Stderr, _colorWarn, _colorError}
	_flag    = log.Ltime | log.Lshortfile
	_level   = Info
	T        = log.New(_workers[Trace], "[T] ", _flag)
	D        = log.New(_workers[Debug], "[D] ", _flag)
	I        = log.New(_workers[Info], "[I] ", _flag)
	W        = log.New(_workers[Warn], "[W] ", _flag)
	E        = log.New(_workers[Error], "[E] ", _flag)
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
	for i := 0; i < len(_writers); i++ {
		_writers[i] = writer
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
	_writers[Trace] = writer

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
	_writers[Debug] = writer

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
	_writers[Info] = writer

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
	_writers[Warn] = writer

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
	_writers[Error] = writer

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
	copy(_workers[:], _writers[:])

	for i := Trace; i < level; i++ {
		_workers[i] = _passer
	}

	setLevel()

	_level = level
}

// GetLevel 获取日志级别.
func GetLevel() Level {
	return _level
}

// nolint: gochecknoinits
func init() {
	setLevel()
}

func setLevel() {
	T = log.New(_workers[Trace], "[T] ", _flag)
	D = log.New(_workers[Debug], "[D] ", _flag)
	I = log.New(_workers[Info], "[I] ", _flag)
	W = log.New(_workers[Warn], "[W] ", _flag)
	E = log.New(_workers[Error], "[E] ", _flag)
}
