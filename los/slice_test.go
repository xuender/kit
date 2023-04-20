package los_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/los"
)

func TestPull(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	elems := []int{1, 2, 2, 4}

	ass.Equal([]int{1, 4}, los.Pull(elems, 2))
	ass.Equal([]int{1, 2, 2, 4}, elems)

	ass.Equal([]int{1, 2, 2, 4}, los.Pull(elems))
	ass.Equal([]int{1, 2, 2, 4}, elems)
}

func TestPullAt(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	elems := []int{1, 2, 2, 4}

	ass.Equal([]int{1, 4}, los.PullAt(elems, 1, 2))
	ass.Equal([]int{1, 2, 2, 4}, elems)

	ass.Equal([]int{1, 2, 2, 4}, los.PullAt(elems))
	ass.Equal([]int{1, 2, 2, 4}, elems)

	ass.Equal([]int{1, 2, 2}, los.PullAt(elems, 3))
	ass.Equal([]int{1, 2, 2, 4}, elems)

	ass.Equal([]int{2, 2, 4}, los.PullAt(elems, 0, 0))
	ass.Equal([]int{1, 2, 2, 4}, elems)

	ass.Equal([]int{4}, los.PullAt(elems, 0, 2, 1))
	ass.Equal([]int{1, 2, 2, 4}, elems)

	ass.Equal([]int{1, 2}, los.PullAt(elems, 3, 1))
	ass.Equal([]int{1, 2, 2, 4}, elems)

	ass.Equal([]int{1, 2, 2, 4}, los.PullAt(elems, -1, 9))
	ass.Equal([]int{1, 2, 2, 4}, elems)

	ass.Equal([]int{1, 6, 7}, los.PullAt([]int{1, 2, 3, 4, 5, 6, 7}, 3, 4, 1, 2))
}
