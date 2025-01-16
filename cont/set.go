package cont

import (
	"iter"
	"maps"
)

// Set represents a collection of unique elements.
// It provides methods to add elements, check for existence, and convert to slices or sequences.
type Set[V comparable] map[V]struct{}

// NewSet creates a new Set with the provided elements.
func NewSet[V comparable](elems ...V) Set[V] {
	set := make(Set[V], len(elems))

	return set.Add(elems...)
}

// Add adds the specified elements to the set.
func (p Set[V]) Add(elems ...V) Set[V] {
	for _, elem := range elems {
		p[elem] = StructNone
	}

	return p
}

// AddSet adds all elements from the provided sets to the set.
func (p Set[V]) AddSet(set ...Set[V]) Set[V] {
	for _, set := range set {
		for elem := range set {
			p[elem] = StructNone
		}
	}

	return p
}

// Has checks if the set contains the specified element.
func (p Set[V]) Has(elem V) bool {
	_, has := p[elem]

	return has
}

// Slice returns a slice of all elements in the set.
func (p Set[V]) Slice() []V {
	elems := make([]V, 0, len(p))

	for elem := range p {
		elems = append(elems, elem)
	}

	return elems
}

// Values returns an iterator sequence of all elements in the set.
func (p Set[V]) Values() iter.Seq[V] {
	return maps.Keys(p)
}

func (p Set[V]) Equal(other Set[V]) bool {
	return maps.Equal(p, other)
}
