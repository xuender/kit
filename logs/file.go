package logs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/xuender/kit/oss"
)

// nolint: gochecknoglobals
var _files = sync.Map{}

// CloseFile 关闭指定日志文件.
func CloseFile(file *os.File) error {
	isEmpty := false
	_ = file.Sync()

	info, err := file.Stat()
	if err == nil {
		isEmpty = info.Size() == 0
	}

	if err := file.Close(); err != nil {
		return err
	}

	if isEmpty {
		if err := os.Remove(file.Name()); err != nil {
			return err
		}
	}

	return nil
}

// Close 关闭日志文件.
func Close() error {
	_files.Range(func(key, value any) bool {
		if file, ok := value.(*os.File); ok {
			if err := CloseFile(file); err != nil {
				return false
			}
		}

		return true
	})

	return nil
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

	fil, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, oss.DefaultFileMode)
	if err != nil {
		return nil, err
	}

	link := filepath.Join(path, name)
	// 关闭旧日志

	if old, has := _files.Load(link); has {
		if file, ok := old.(*os.File); ok {
			_ = CloseFile(file)
		}
	}

	_files.Store(link, fil)
	_ = os.Remove(link)
	_ = os.Symlink(log, link)

	return fil, err
}
