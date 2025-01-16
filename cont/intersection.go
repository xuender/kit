package cont

import (
	"iter"
	"slices"

	"github.com/xuender/kit/v2/types"
)

// Intersection returns a sequence that contains the intersection of multiple slices.
// It yields elements that are present in all provided slices, ensuring each element appears only once.
func Intersection[S ~[]E, E comparable](items ...S) iter.Seq[E] {
	return func(yield func(E) bool) {
		if len(items) == 0 {
			return
		}

		set := NewSet[E]()

		for _, item := range items[0] {
			if set.Has(item) {
				continue
			}

			set.Add(item)

			has := true

			for _, slice := range items[1:] {
				if !slices.Contains(slice, item) {
					has = false

					break
				}
			}

			if has && !yield(item) {
				return
			}
		}
	}
}

// IntersectionFunc returns a sequence that contains the intersection of multiple slices using equality function.
func IntersectionFunc[S ~[]E, E any](equal types.Equaler[E], items ...S) iter.Seq[E] {
	return func(yield func(E) bool) {
		if len(items) == 0 {
			return
		}

		intersection := []E{}

		for _, item := range items[0] {
			if slices.ContainsFunc(intersection, func(val E) bool { return equal(val, item) }) {
				continue
			}

			pass := false

			for _, slice := range items[1:] {
				if !slices.ContainsFunc(slice, func(val E) bool { return equal(val, item) }) {
					pass = true

					break
				}
			}

			if pass {
				continue
			}

			if !yield(item) {
				return
			}

			intersection = append(intersection, item)
		}
	}
}
