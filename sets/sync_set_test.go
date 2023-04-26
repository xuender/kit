package sets_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/sets"
)

func TestSyncSet_Iteration(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	set := sets.NewSyncSet(1, 2, 3, 4, 5)

	ass.Error(set.Iteration(func(i int) error {
		if i > 4 {
			return assert.AnError
		}

		ass.LessOrEqual(i, 4)

		return nil
	}))

	ass.Nil(set.Iteration(func(i int) error {
		ass.LessOrEqual(i, 5)

		return nil
	}))
}
