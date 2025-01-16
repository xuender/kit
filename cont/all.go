package cont

import "github.com/xuender/kit/v2/types"

// All returns true if all elements in the slice satisfy every provided checker function.
// It returns false if any element fails to satisfy any of the checkers.
func All[S ~[]E, E any](items S, checkers ...types.Checker[E]) bool {
	for _, checker := range checkers {
		for _, item := range items {
			if !checker(item) {
				return false
			}
		}
	}

	return true
}
