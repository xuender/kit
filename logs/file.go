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
	path, err := oss.Abs(path)
	if err != nil {
		return nil, err
	}

	_ = os.MkdirAll(path, oss.DefaultDirFileMod)

	link := filepath.Join(path, name)

	ext := filepath.Ext(name)
	suffix := time.Now().Format("06010215")
	log := fmt.Sprintf("%s-%s%s", name[:len(name)-len(ext)], suffix, ext)

	if ext == "" {
		log = fmt.Sprintf("%s-%s", name, suffix)
	}

	file, err := oss.Abs(filepath.Join(path, log))
	if err != nil {
		return nil, err
	}

	writer, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, oss.DefaultFileMode)
	if err != nil {
		return nil, err
	}

	_closers = append(_closers, writer)
	_ = os.Remove(link)
	_ = os.Symlink(log, link)

	return writer, err
}
