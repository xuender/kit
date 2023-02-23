package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/base"
)

func TestIf(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal(1, base.If(true, 1, 2))
	ass.Equal(2, base.If(false, 1, 2))
}
