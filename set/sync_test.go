package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kit/set"
)

func TestSync_Iterate(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	req := require.New(t)
	set := set.NewSync(1, 2, 3, 4, 5)

	req.Error(set.Iterate(func(i int) error {
		if i > 4 {
			return assert.AnError
		}

		ass.LessOrEqual(i, 4)

		return nil
	}))

	req.NoError(set.Iterate(func(i int) error {
		ass.LessOrEqual(i, 5)

		return nil
	}))
}
