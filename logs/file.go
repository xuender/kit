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
var _closers = map[string]io.Closer{}

// Close 关闭日志文件.
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

	ext := filepath.Ext(name)
	suffix := time.Now().Format("06010215")
	log := fmt.Sprintf("%s-%s%s", name[:len(name)-len(ext)], suffix, ext)

	file, _ := oss.Abs(filepath.Join(path, log))

	writer, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, oss.DefaultFileMode)
	if err != nil {
		return nil, err
	}

	link := filepath.Join(path, name)
	// 关闭旧日志
	if old, has := _closers[link]; has {
		old.Close()
	}

	_closers[link] = writer
	_ = os.Remove(link)
	_ = os.Symlink(log, link)

	return writer, err
}
