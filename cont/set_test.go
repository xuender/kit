package cont_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/v2/cont"
)

func TestSet_AddAll(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	nums := cont.NewSet(1, 2, 3)

	nums.AddSet(cont.NewSet(3, 4, 5))
	ass.Len(nums, 5)
}

func TestSet_Values(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	nums := cont.NewSet(1)

	for num := range nums.Values() {
		ass.Equal(1, num)
	}
}

func TestSet_Equal(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	set1 := cont.NewSet[int](1, 2, 3)
	set2 := cont.NewSet[int](3, 2, 1)

	ass.True(set1.Equal(set2))

	set2.Add(3)
	ass.True(set1.Equal(set2))

	set2.Add(4)
	ass.False(set1.Equal(set2))
}
