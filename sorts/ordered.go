package sorts

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type ordered[O constraints.Ordered] []O

func (p ordered[O]) Len() int           { return len(p) }
func (p ordered[O]) Less(i, j int) bool { return p[i] < p[j] }
func (p ordered[O]) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Ordered[O constraints.Ordered](elems []O) {
	sort.Sort(ordered[O](elems))
}
