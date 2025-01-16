package cont

import (
	"iter"
	"slices"

	"github.com/xuender/kit/v2/types"
)

// Union returns a sequence that contains the union of multiple slices.
// It eliminates duplicate elements using a set, ensuring each element appears only once.
func Union[S ~[]E, E comparable](items ...S) iter.Seq[E] {
	return func(yield func(E) bool) {
		if len(items) == 0 {
			return
		}

		set := NewSet[E]()

		for _, slice := range items {
			for _, item := range slice {
				if set.Has(item) {
					continue
				}

				set.Add(item)

				if !yield(item) {
					return
				}
			}
		}
	}
}

// UnionFunc returns a sequence that contains the union of multiple slices using a custom equality function.
func UnionFunc[S ~[]E, E any](equal types.Equaler[E], items ...S) iter.Seq[E] {
	return func(yield func(E) bool) {
		if len(items) == 0 {
			return
		}

		union := []E{}

		for _, slice := range items {
			for _, item := range slice {
				if slices.ContainsFunc(union, func(val E) bool {
					return equal(val, item)
				}) {
					continue
				}

				union = append(union, item)

				if !yield(item) {
					return
				}
			}
		}
	}
}
