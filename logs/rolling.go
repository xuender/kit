package logs

import (
	"io"

	"github.com/arthurkiller/rollingwriter"
	"github.com/xuender/kit/base"
)

type rolling struct {
	writer io.WriteCloser
}

func newRolling(path, name string) (*rolling, error) {
	cfg := &rollingwriter.Config{
		TimeTagFormat:          "060102150405",
		LogPath:                path,
		FileName:               name,
		MaxRemain:              base.TwoHundredFiftySix,
		RollingPolicy:          rollingwriter.VolumeRolling,
		RollingTimePattern:     "* * * * * *",
		RollingVolumeSize:      "100M",
		WriterMode:             "lock",
		BufferWriterThershould: base.Eight * base.OneHundredTwentyEight * base.OneHundredTwentyEight,
		// Compress:               true,
	}
	writer, err := rollingwriter.NewWriterFromConfig(cfg)

	return &rolling{writer: writer}, err
}

func (p *rolling) Write(data []byte) (int, error) {
	return p.writer.Write(data)
}

func (p *rolling) Close() error {
	return p.writer.Close()
}
