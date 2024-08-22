package los

import (
	"sort"

	"github.com/samber/lo"
)

// Pull 删除切片指定成员.
func Pull[V comparable](collection []V, elements ...V) []V {
	if len(elements) == 0 {
		return collection
	}

	return Remove(collection, func(item V, _ int) bool { return lo.Contains(elements, item) })
}

// PullAt 删除切片指定位置.
func PullAt[S ~[]V, V any](collection S, indices ...int) S {
	indices = lo.Union(lo.Filter(indices, func(item, _ int) bool { return item >= 0 && item < len(collection) }))

	if len(indices) > 1 && !lo.IsSorted(indices) {
		sort.Ints(indices)
	}

	ret := make([]V, len(collection)-len(indices))
	start, left := 0, 0

	for _, right := range indices {
		start += copy(ret[start:], collection[left:right])
		left = right + 1
	}

	copy(ret[start:], collection[left:])

	return ret
}

// Remove 根据断言删除.
func Remove[V comparable](collection []V, predicate func(V, int) bool) []V {
	return lo.Filter(collection, func(item V, index int) bool { return !predicate(item, index) })
}
