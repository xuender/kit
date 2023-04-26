package sets

import "sync"

// SyncSet 线程安全Set.
type SyncSet[V comparable] struct {
	mutex sync.RWMutex
	data  map[V]struct{}
}

// NewSyncSet 新建线程安全Set.
func NewSyncSet[V comparable](elems ...V) *SyncSet[V] {
	set := &SyncSet[V]{mutex: sync.RWMutex{}, data: map[V]struct{}{}}

	return set.Add(elems...)
}

// Len 集合长度.
func (p *SyncSet[V]) Len() int {
	return len(p.data)
}

// Add 增加元素.
func (p *SyncSet[V]) Add(elems ...V) *SyncSet[V] {
	p.mutex.Lock()

	for _, elem := range elems {
		p.data[elem] = struct{}{}
	}

	p.mutex.Unlock()

	return p
}

// Delete 删除元素.
func (p *SyncSet[V]) Delete(elems ...V) *SyncSet[V] {
	p.mutex.Lock()

	for _, elem := range elems {
		delete(p.data, elem)
	}

	p.mutex.Unlock()

	return p
}

// Has 包含.
func (p *SyncSet[V]) Has(elem V) bool {
	p.mutex.RLock()

	_, has := p.data[elem]

	p.mutex.RUnlock()

	return has
}

// Slice 转换成切片.
func (p *SyncSet[V]) Slice() []V {
	p.mutex.RLock()

	elems := make([]V, 0, len(p.data))
	for elem := range p.data {
		elems = append(elems, elem)
	}

	p.mutex.RUnlock()

	return elems
}

// Iteration 迭代.
func (p *SyncSet[V]) Iteration(yield func(V) error) error {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for elem := range p.data {
		if err := yield(elem); err != nil {
			return err
		}
	}

	return nil
}
