package logs

import (
	"os"

	"github.com/xuender/kit/base"
)

type color []byte

// nolint: gochecknoglobals
var (
	_reset            = []byte{27, 91, 48, 109}
	_colorError color = []byte{27, 91, 51, 49, 109}
	_colorWarn  color = []byte{27, 91, 51, 51, 109}
	_colorDebug color = []byte{27, 91, 51, 52, 109}
	_colorTrace color = []byte{27, 91, 51, 53, 109}
)

func (p color) Write(data []byte) (int, error) {
	ret := make([]byte, len(data)+base.Nine)
	copy(ret, p)
	copy(ret[5:], data)
	copy(ret[len(data)+5:], _reset)

	return os.Stderr.Write(ret)
}
