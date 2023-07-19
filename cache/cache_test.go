package cache_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/cache"
)

func TestCache_Get(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cac := cache.New[int, int](time.Millisecond*2, cache.NoExpiration)

	defer cac.Close()

	cac.Set(1, 1)

	time.Sleep(time.Millisecond)

	_, found := cac.Get(1)
	ass.True(found)

	time.Sleep(time.Millisecond)

	_, found = cac.Get(1)
	ass.True(found)

	time.Sleep(time.Millisecond)

	_, found = cac.Get(1)
	ass.True(found)

	time.Sleep(time.Millisecond * 2)

	_, found = cac.Get(1)
	ass.False(found)
}

func TestCache_GetNoExtension(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cac := cache.New[int, int](time.Millisecond, cache.NoExpiration)

	defer cac.Close()

	cac.Set(1, 1)
	time.Sleep(time.Millisecond)

	_, has := cac.GetNoExtension(1)
	ass.False(has)
}

func TestCache_Iterate(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	cac := cache.New[int, int](cache.DefaultExpiration, cache.NoExpiration)

	defer cac.Close()

	for i := 0; i < 10; i++ {
		cac.Set(i, i)
	}

	cac.SetDuration(100, 100, time.Millisecond)
	time.Sleep(time.Millisecond)

	ass.Nil(cac.Iterate(func(key, value int) error {
		ass.Equal(key, value)

		return nil
	}))

	ass.Error(cac.Iterate(func(key, value int) error {
		ass.Equal(key, value)

		if key > 3 {
			return assert.AnError
		}

		return nil
	}))
}
