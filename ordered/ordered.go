package ordered

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Ordered[O constraints.Ordered] []O

func (p Ordered[O]) Len() int           { return len(p) }
func (p Ordered[O]) Less(i, j int) bool { return p[i] < p[j] }
func (p Ordered[O]) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort ordered.
func Sort[O constraints.Ordered](elems []O) {
	sort.Sort(Ordered[O](elems))
}
