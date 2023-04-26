package sets

// MapSet 基于map的Set.
type MapSet[V comparable] map[V]struct{}

// NewMapSet 新建MapSet.
func NewMapSet[V comparable](elems ...V) MapSet[V] {
	set := make(MapSet[V], len(elems))

	return set.Add(elems...)
}

// Add 增加元素.
func (p MapSet[V]) Add(elems ...V) MapSet[V] {
	for _, elem := range elems {
		p[elem] = struct{}{}
	}

	return p
}

// AddSet 增加集合.
func (p MapSet[V]) AddSet(sets ...MapSet[V]) MapSet[V] {
	for _, set := range sets {
		for elem := range set {
			p[elem] = struct{}{}
		}
	}

	return p
}

// Has 包含.
func (p MapSet[V]) Has(elem V) bool {
	_, has := p[elem]

	return has
}

// Slice 转换切片.
func (p MapSet[V]) Slice() []V {
	elems := make([]V, 0, len(p))

	for elem := range p {
		elems = append(elems, elem)
	}

	return elems
}
