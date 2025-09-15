package ordered

import (
	"cmp"
	"sort"
)

type Ordered[T cmp.Ordered] []T

func (p Ordered[T]) Len() int           { return len(p) }
func (p Ordered[T]) Less(i, j int) bool { return p[i] < p[j] }
func (p Ordered[T]) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort ordered.
func Sort[T cmp.Ordered](elems []T) {
	sort.Sort(Ordered[T](elems))
}
