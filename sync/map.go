package sync

import (
	"sync"

	"github.com/samber/lo"
)

type Map[K comparable, V any] struct{ sync.Map }

func NewMap[K comparable, V any](items ...lo.Tuple2[K, V]) *Map[K, V] {
	ret := &Map[K, V]{}

	for _, item := range items {
		ret.Store(item.A, item.B)
	}

	return ret
}

// func (p *Map[K, V]) CompareAndDelete(key K, old V) bool {
// 	return p.Map.CompareAndDelete(key, old)
// }

// func (p *Map[K, V]) CompareAndSwap(key K, oldVal, newVal V) bool {
// 	return p.Map.CompareAndSwap(key, oldVal, newVal)
// }

func (p *Map[K, V]) Delete(key K) { p.Map.Delete(key) }

func (p *Map[K, V]) Load(key K) (V, bool) {
	var zero V

	value, has := p.Map.Load(key)
	if !has {
		return zero, false
	}

	ret, _ := value.(V)

	return ret, has
}

func (p *Map[K, V]) LoadAndDelete(key K) (V, bool) {
	val, has := p.Map.LoadAndDelete(key)
	ret, _ := val.(V)

	return ret, has
}

func (p *Map[K, V]) LoadOrCreate(key K, create func() V) (V, bool) {
	if val, has := p.Load(key); has {
		return val, has
	}

	return p.LoadOrStore(key, create())
}

func (p *Map[K, V]) LoadOrStore(key K, value V) (V, bool) {
	val, has := p.Map.LoadOrStore(key, value)
	ret, _ := val.(V)

	return ret, has
}

func (p *Map[K, V]) Range(call func(key K, value V) bool) {
	p.Map.Range(func(key, value any) bool {
		rkey, _ := key.(K)
		rvalue, _ := value.(V)

		return call(rkey, rvalue)
	})
}

func (p *Map[K, V]) Store(key K, value V) { p.Map.Store(key, value) }

// func (p *Map[K, V]) Swap(key K, value V) (V, bool) {
// 	val, has := p.Map.Swap(key, value)
// 	ret, _ := val.(V)

// 	return ret, has
// }
