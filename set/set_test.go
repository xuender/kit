package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/set"
)

func TestSet_AddAll(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	nums := set.NewSet(1, 2, 3)

	nums.AddSet(set.NewSet(3, 4, 5))
	ass.Len(nums, 5)
}
