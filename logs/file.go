package logs

import (
	"io"
	"log"
)

// nolint: gochecknoglobals
var _closers = []io.Closer{}

func LogFile(path, name string) (io.Writer, error) {
	_flag = log.Ltime | log.Ldate | log.Lshortfile

	writer, err := newRolling(path, name)
	if err == nil {
		_closers = append(_closers, writer)
	}

	return writer, err
}

func Close() {
	for _, closer := range _closers {
		closer.Close()
	}
}
