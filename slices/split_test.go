package slices_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/slices"
)

func TestSplitStr(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	ass.Equal("a", slices.SplitStr("a-b_c", '-', '_')[0])
	ass.Equal("b", slices.SplitStr("a-b_c", '-', '_')[1])
	ass.Equal("c", slices.SplitStr("a-b_c", '-', '_')[2])

	ass.Equal("c", slices.SplitStr("a-_c", '-', '_')[2])
	ass.Equal("", slices.SplitStr("a-_c", '-', '_')[1])
}
