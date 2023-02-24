package logs

import (
	"io"

	"github.com/arthurkiller/rollingwriter"
	"github.com/xuender/kit/base"
)

func NewRolling(path, name string) (io.WriteCloser, error) {
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

	return rollingwriter.NewWriterFromConfig(cfg)
}
