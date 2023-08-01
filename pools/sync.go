package pools

import "sync"

type SyncPool[T any] struct {
	sync.Pool
	init func(T)
}

func NewSyncPool[T any](newFunc func() T, initFunc func(T)) *SyncPool[T] {
	return &SyncPool[T]{sync.Pool{New: func() any { return newFunc() }}, initFunc}
}

// nolint: forcetypeassert
func (p *SyncPool[T]) Get() T {
	ret := p.Pool.Get().(T)

	if p.init != nil {
		p.init(ret)
	}

	return ret
}

func (p *SyncPool[T]) Put(elem T) {
	p.Pool.Put(elem)
}
