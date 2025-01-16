package cont

import "github.com/xuender/kit/v2/types"

// Any returns true if any element in the slice satisfies at least one of the provided checker functions.
// It returns false if no elements satisfy any of the checkers.
func Any[S ~[]E, E any](items S, checkers ...types.Checker[E]) bool {
	for _, checker := range checkers {
		for _, item := range items {
			if checker(item) {
				return true
			}
		}
	}

	return false
}
