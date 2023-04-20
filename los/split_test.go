package los_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/los"
)

func TestSplitStr(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	ass.Equal("a", los.SplitStr("a-b_c", '-', '_')[0])
	ass.Equal("b", los.SplitStr("a-b_c", '-', '_')[1])
	ass.Equal("c", los.SplitStr("a-b_c", '-', '_')[2])

	ass.Equal("c", los.SplitStr("a-_c", '-', '_')[2])
	ass.Equal("", los.SplitStr("a-_c", '-', '_')[1])
}
