package sync_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/sync"
)

func TestMap_Load(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap[int, string]()
	cmap.Store(3, "3")
	cmap.Store(2, "2")

	val, has := cmap.Load(1)
	ass.False(has)
	ass.Equal("", val)

	val, has = cmap.Load(2)
	ass.True(has)
	ass.Equal("2", val)

	val, has = cmap.Load(3)
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

	cmap.Store(3, &ts{num: 4})

	val, has := cmap.Load(3)
	ass.True(has)
	ass.Equal(4, val.num)
}

func TestMap_map(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap[int, *sync.Map[int, int]]()

	cmap.Store(3, sync.NewMap[int, int]())

	val, has := cmap.Load(3)
	ass.True(has)
	val.Store(9, 9)

	val2, _ := val.Load(9)
	ass.True(has)
	ass.Equal(9, val2)
}

func TestMap_LoadOrStore(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap[int, int]()

	val, has := cmap.LoadOrStore(1, 1)
	ass.False(has)
	ass.Equal(1, val)

	val, has = cmap.LoadOrStore(1, 7)
	ass.True(has)
	ass.Equal(1, val)
}

func TestMap_LoadOrCreate(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap[int, int]()

	create := func() int { return 1 }

	val, has := cmap.LoadOrCreate(1, create)
	ass.False(has)
	ass.Equal(1, val)

	val, has = cmap.LoadOrCreate(1, create)
	ass.True(has)
	ass.Equal(1, val)
}

func TestMap_LoadOrDelete(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap[int, int]()

	cmap.Store(1, 1)

	val, has := cmap.LoadAndDelete(1)
	ass.True(has)
	ass.Equal(1, val)

	_, has = cmap.Load(1)
	ass.False(has)
}

func TestMap_NewMap(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap(lo.T2(1, "1"), lo.T2(2, "2"))
	val, has := cmap.Load(1)

	ass.True(has)
	ass.Equal("1", val)

	cmap.Delete(1)

	_, has = cmap.Load(1)

	ass.False(has)
}

func TestMap_Range(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap(lo.T2(1, "1"))

	cmap.Range(func(key int, value string) bool {
		ass.Equal(1, key)
		ass.Equal("1", value)

		return true
	})
}

func TestMap_Swap(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cmap := sync.NewMap(lo.T2(1, "1"), lo.T2(2, "2"))

	cmap.Swap(1, "3")
	cmap.CompareAndSwap(1, "3", "9")

	val, _ := cmap.Load(1)
	ass.Equal("9", val)

	ass.False(cmap.CompareAndDelete(1, "3"))
	ass.True(cmap.CompareAndDelete(1, "9"))
}
