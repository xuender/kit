package cont

import (
	"github.com/xuender/kit/v2/types"
)

// LastIndex returns the last index of an item in the slice. If the item is not found, it returns -1.
// It searches from the end of the slice to the beginning.
func LastIndex[S ~[]E, E comparable](items S, item E) int {
	var idx int

	for idx = len(items) - 1; idx >= 0; idx-- {
		if items[idx] == item {
			return idx
		}
	}

	return idx
}

// LastIndexFunc returns the last index of an item that satisfies the checker function in the slice.
// If no item satisfies the checker, it returns -1. It searches from the end of the slice to the beginning.
func LastIndexFunc[S ~[]E, E any](items S, checker types.Checker[E]) int {
	var idx int

	for idx = len(items) - 1; idx >= 0; idx-- {
		if checker(items[idx]) {
			return idx
		}
	}

	return idx
}
