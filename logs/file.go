package logs

import (
	"io"
)

// nolint: gochecknoglobals
var _closers = []io.Closer{}

func LogFile(path, name string) (io.Writer, error) {
	writer, err := NewRolling(path, name)
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
