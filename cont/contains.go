package cont

import (
	"github.com/xuender/kit/v2/types"
)

// Contains checks if a slice contains an element based on a custom equality function.
// It returns true if the specified item is found in the slice.
func Contains[S ~[]E, E any](equal types.Equaler[E], items S, item E) bool {
	if len(items) == 0 {
		return false
	}

	for _, it := range items {
		if equal(item, it) {
			return true
		}
	}

	return false
}
