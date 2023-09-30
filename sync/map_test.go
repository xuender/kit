package sync_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/sync"
)

func TestMap_Get(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap[int, string]()
	cmap.Set(3, "3")
	cmap.Set(2, "2")

	val, has := cmap.Get(1)
	ass.False(has)
	ass.Equal("", val)

	val, has = cmap.Get(2)
	ass.True(has)
	ass.Equal("2", val)

	val, has = cmap.Get(3)
	ass.True(has)
	ass.Equal("3", val)
}

type ts struct {
	num int
}

func TestMap_status(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap[int, *ts]()

	cmap.Set(3, &ts{num: 4})

	val, has := cmap.Get(3)
	ass.True(has)
	ass.Equal(4, val.num)
}

func TestMap_map(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap[int, *sync.Map[int, int]]()

	cmap.Set(3, sync.NewMap[int, int]())

	val, has := cmap.Get(3)
	ass.True(has)
	val.Set(9, 9)

	val2, _ := val.Get(9)
	ass.True(has)
	ass.Equal(9, val2)
}

func TestMap_GetOrSet(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap[int, int]()

	val, has := cmap.GetOrSet(1, 1)
	ass.False(has)
	ass.Equal(1, val)

	val, has = cmap.GetOrSet(1, 7)
	ass.True(has)
	ass.Equal(1, val)
}

func TestMap_GetOrDelete(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap[int, int]()

	cmap.Set(1, 1)

	val, has := cmap.GetOrDelete(1)
	ass.True(has)
	ass.Equal(1, val)

	_, has = cmap.Get(1)
	ass.False(has)
}

func TestMap_NewMap(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap(lo.T2(1, "1"), lo.T2(2, "2"))
	val, has := cmap.Get(1)

	ass.True(has)
	ass.Equal("1", val)
}
