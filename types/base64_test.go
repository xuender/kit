package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/types"
)

func TestNumToB64(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	for i := 0; i < 51; i++ {
		for f := 0; f < 10; f++ {
			num := 1 << i & f
			ass.Equal(num, types.B64ToNum[int](types.NumToB64(num)))
		}
	}
}
