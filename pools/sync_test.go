package pools_test

import (
	"sync"
	"testing"

	"github.com/xuender/kit/pools"
)

func BenchmarkNew(b *testing.B) {
	var elem *data

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			elem = &data{}
			elem.data = j
		}
	}
}

// nolint: forcetypeassert
func BenchmarkSync(b *testing.B) {
	pool := sync.Pool{New: func() any { return &data{1} }}

	var elem *data

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			elem = pool.Get().(*data)
			elem.data = j
			pool.Put(elem)
		}
	}
}

func BenchmarkSyncPool(b *testing.B) {
	pool := pools.NewSyncPool(func() *data { return &data{1} }, func(d *data) { d.data = 1 })

	var elem *data

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			elem = pool.Get()
			elem.data = j
			pool.Put(elem)
		}
	}
}
