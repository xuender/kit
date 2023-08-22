package ordered

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Ordered[T constraints.Ordered] []T

func (p Ordered[T]) Len() int           { return len(p) }
func (p Ordered[T]) Less(i, j int) bool { return p[i] < p[j] }
func (p Ordered[T]) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort ordered.
func Sort[T constraints.Ordered](elems []T) {
	sort.Sort(Ordered[T](elems))
}
