package cont

import (
	"iter"
	"slices"

	"github.com/xuender/kit/v2/types"
)

// Difference returns a sequence that src slice which are not present in any of the provided slices.
func Difference[S ~[]E, E comparable](src S, items ...S) iter.Seq[E] {
	if len(items) == 0 {
		return slices.Values(src)
	}

	return func(yield func(E) bool) {
		for _, item := range src {
			if Any(items, func(slice S) bool { return slices.Contains(slice, item) }) {
				continue
			}

			if !yield(item) {
				return
			}
		}
	}
}

// DifferenceFunc returns a sequence that src slice which are not present in any of the provided slices.
func DifferenceFunc[S ~[]E, E any](equal types.Equaler[E], src S, items ...S) iter.Seq[E] {
	if len(items) == 0 {
		return slices.Values(src)
	}

	return func(yield func(E) bool) {
		for _, item := range src {
			if Any(items, func(slice S) bool { return Contains(equal, slice, item) }) {
				continue
			}

			if !yield(item) {
				return
			}
		}
	}
}
