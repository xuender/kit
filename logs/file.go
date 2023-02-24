package logs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/xuender/kit/oss"
)

// nolint: gochecknoglobals
var (
	_closers = []io.Closer{}
)

func Close() {
	for _, closer := range _closers {
		closer.Close()
	}
}

// File 生成软连接文件.
func File(path, name string) (io.Writer, error) {
	link, err := oss.Abs(filepath.Join(path, name))
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(name)
	suffix := time.Now().Format("06010215")
	log := fmt.Sprintf("%s-%s%s", name[:len(ext)], suffix, ext)

	file, err := oss.Abs(filepath.Join(path, log))
	if err != nil {
		return nil, err
	}

	writer, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, err
	}

	_closers = append(_closers, writer)
	_ = os.Remove(link)
	_ = os.Symlink(log, link)

	return writer, err
}
