package set

import (
	"sync"

	"github.com/xuender/kit/base"
)

// Sync 线程安全Set.
type Sync[V comparable] struct {
	mutex sync.RWMutex
	data  map[V]struct{}
}

// NewSync 新建线程安全Set.
func NewSync[V comparable](elems ...V) *Sync[V] {
	set := &Sync[V]{sync.RWMutex{}, map[V]struct{}{}}

	return set.Add(elems...)
}

// Len 集合长度.
func (p *Sync[V]) Len() int {
	return len(p.data)
}

// Add 增加元素.
func (p *Sync[V]) Add(elems ...V) *Sync[V] {
	p.mutex.Lock()

	for _, elem := range elems {
		p.data[elem] = base.None
	}

	p.mutex.Unlock()

	return p
}

// Delete 删除元素.
func (p *Sync[V]) Delete(elems ...V) *Sync[V] {
	p.mutex.Lock()

	for _, elem := range elems {
		delete(p.data, elem)
	}

	p.mutex.Unlock()

	return p
}

// Has 包含.
func (p *Sync[V]) Has(elem V) bool {
	p.mutex.RLock()

	_, has := p.data[elem]

	p.mutex.RUnlock()

	return has
}

// Slice 转换成切片.
func (p *Sync[V]) Slice() []V {
	p.mutex.RLock()

	elems := make([]V, 0, len(p.data))
	for elem := range p.data {
		elems = append(elems, elem)
	}

	p.mutex.RUnlock()

	return elems
}

// Iterate 迭代.
func (p *Sync[V]) Iterate(yield func(V) error) error {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for elem := range p.data {
		if err := yield(elem); err != nil {
			return err
		}
	}

	return nil
}
