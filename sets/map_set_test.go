package sets_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/sets"
)

func TestMapSet_AddAll(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	set := sets.NewMapSet(1, 2, 3)

	set.AddSet(sets.NewMapSet(3, 4, 5))
	ass.Equal(5, len(set))
}
