package sync

import (
	"sync"

	"github.com/samber/lo"
)

type Map[K, V any] struct {
	sync.Map
}

func NewMap[K, V any](items ...lo.Tuple2[K, V]) *Map[K, V] {
	ret := &Map[K, V]{}

	for _, item := range items {
		ret.Set(item.A, item.B)
	}

	return ret
}

func (p *Map[K, V]) Get(key K) (V, bool) {
	var zero V

	value, has := p.Load(key)
	if !has {
		return zero, false
	}

	return value.(V), true
}

func (p *Map[K, V]) Set(key K, value V) {
	p.Store(key, value)
}

func (p *Map[K, V]) GetOrSet(key K, value V) (V, bool) {
	val, has := p.LoadOrStore(key, value)

	return val.(V), has
}

func (p *Map[K, V]) GetOrDelete(key K) (V, bool) {
	val, has := p.LoadAndDelete(key)

	return val.(V), has
}
