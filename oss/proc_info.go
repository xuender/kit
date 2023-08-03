package oss

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/xuender/kit/times"
)

// nolint: gochecknoglobals
var (
	// Version 版本号.
	Version = ""
	// BuildTime 编译时间.
	BuildTime = ""
)

type ProcInfo struct {
	Name      string        `json:"name"`
	Dir       string        `json:"dir"`
	Pid       int           `json:"pid"`
	Version   string        `json:"version"`
	BuildTime string        `json:"buildTime"`
	StartTime time.Time     `json:"startTime"`
	RunTime   time.Duration `json:"runTime"`
}

func NewBuildInfo() *ProcInfo {
	path := os.Args[0]
	name := filepath.Base(path)
	dir := filepath.Dir(path)

	return &ProcInfo{
		name,
		dir,
		os.Getpid(),
		Version,
		BuildTime,
		time.Now(),
		0,
	}
}

func (p *ProcInfo) String() string {
	p.RunTime = time.Since(p.StartTime)

	return fmt.Sprintf(`Name: %s
Pid: %d
Dir: %s
Version: %s
BuildTime: %s
StartTime: %s
RunTime: %s
`,
		p.Name,
		p.Pid,
		p.Dir,
		p.Version,
		p.BuildTime,
		p.StartTime,
		times.Duration(p.RunTime),
	)
}
