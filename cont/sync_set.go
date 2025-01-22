package cont

import (
	"iter"
	"maps"
	"sync"
)

type SyncSet[V comparable] struct {
	items map[V]struct{}
	lock  sync.RWMutex
}

func NewSyncSet[V comparable](elems ...V) *SyncSet[V] {
	items := make(map[V]struct{}, len(elems))
	for _, elem := range elems {
		items[elem] = StructNone
	}

	return &SyncSet[V]{
		items: items,
	}
}

func (p *SyncSet[V]) Add(elems ...V) *SyncSet[V] {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, elem := range elems {
		p.items[elem] = StructNone
	}

	return p
}

func (p *SyncSet[V]) Has(elem V) bool {
	p.lock.RLock()
	defer p.lock.RUnlock()

	_, has := p.items[elem]

	return has
}

func (p *SyncSet[V]) Len() int {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return len(p.items)
}

func (p *SyncSet[V]) Del(elem V) *SyncSet[V] {
	p.lock.Lock()
	defer p.lock.Unlock()

	delete(p.items, elem)

	return p
}

func (p *SyncSet[V]) Slice() []V {
	p.lock.RLock()
	defer p.lock.RUnlock()

	elems := make([]V, 0, len(p.items))

	for elem := range p.items {
		elems = append(elems, elem)
	}

	return elems
}

// Values returns an iterator sequence of all elements in the set.
func (p *SyncSet[V]) Values() iter.Seq[V] {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return maps.Keys(p.items)
}
