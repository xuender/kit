package cache_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kit/cache"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

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
	res := require.New(t)
	cac := cache.New[int, int](cache.DefaultExpiration, cache.NoExpiration)

	defer cac.Close()

	for i := range 10 {
		cac.Set(i, i)
	}

	cac.SetDuration(100, 100, time.Millisecond)
	time.Sleep(time.Millisecond)

	res.NoError(cac.Iterate(func(key, value int) error {
		ass.Equal(key, value)

		return nil
	}))

	res.Error(cac.Iterate(func(key, value int) error {
		ass.Equal(key, value)

		if key > 3 {
			return assert.AnError
		}

		return nil
	}))
}
