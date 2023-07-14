package sets

import "github.com/xuender/kit/base"

// Set 基于map的Set.
type Set[V comparable] map[V]struct{}

// NewSet 新建MapSet.
func NewSet[V comparable](elems ...V) Set[V] {
	set := make(Set[V], len(elems))

	return set.Add(elems...)
}

// Add 增加元素.
func (p Set[V]) Add(elems ...V) Set[V] {
	for _, elem := range elems {
		p[elem] = base.None
	}

	return p
}

// AddSet 增加集合.
func (p Set[V]) AddSet(sets ...Set[V]) Set[V] {
	for _, set := range sets {
		for elem := range set {
			p[elem] = base.None
		}
	}

	return p
}

// Has 包含.
func (p Set[V]) Has(elem V) bool {
	_, has := p[elem]

	return has
}

// Slice 转换切片.
func (p Set[V]) Slice() []V {
	elems := make([]V, 0, len(p))

	for elem := range p {
		elems = append(elems, elem)
	}

	return elems
}
