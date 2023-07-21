package counter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/counter"
)

func TestNewCounter(t *testing.T) {
	t.Parallel()

	count := counter.NewCounter[int]()
	for i := 0; i < 1000; i++ {
		count.Inc(1)
	}

	for i := 0; i < 2000; i++ {
		count.Inc(2)
	}

	assert.Equal(t, int64(1000), count.Get(1))
	assert.Equal(t, int64(2000), count.Get(2))
}

func BenchmarkCounter(b *testing.B) {
	count := counter.NewCounter[int]()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		count.Inc(1)
	}
}

// func BenchmarkCounterb(b *testing.B) {
// 	count := syncs.NewCounterb[int]()

// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
//		count.Inc(1)
// 	}
// }

func TestCounter_Keys(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	count := counter.NewCounter[int]()

	for i := 0; i < 100; i++ {
		count.Inc(i)
	}

	ass.Equal(100, len(count.Keys()))
}

func TestCounter_Dec(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	count := counter.NewCounter[int]()

	count.Dec(1)
	ass.Equal(int64(-1), count.Get(1))

	for i := 0; i < 100; i++ {
		count.Inc(1)
	}

	ass.Equal(int64(99), count.Get(1))
	count.Dec(1)
	ass.Equal(int64(98), count.Get(1))
}

func TestCounter_Size(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	count := counter.NewCounter[int]()

	for i := 0; i < 100; i++ {
		count.Inc(i)
	}

	ass.Equal(100, count.Size())
}

func TestCounter_Sum(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	count := counter.NewCounter[int]()

	for i := 0; i < 100; i++ {
		count.Inc(i)
	}

	ass.Equal(int64(100), count.Sum())
	count.Clean()
	ass.Equal(int64(0), count.Sum())
}
