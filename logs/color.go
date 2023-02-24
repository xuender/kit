package logs

import (
	"os"

	"github.com/xuender/kit/ios"
)

type color []byte

// nolint: gochecknoglobals
var (
	_reset            = []byte{27, 91, 48, 109}
	_colorTrace color = []byte{27, 91, 51, 54, 109}
	_colorDebug color = []byte{27, 91, 51, 50, 109}
	_colorWarn  color = []byte{27, 91, 51, 51, 109}
	_colorError color = []byte{27, 91, 51, 49, 109}
)

func (p color) Write(data []byte) (int, error) {
	return ios.Write(os.Stderr, p, data, _reset)
}
