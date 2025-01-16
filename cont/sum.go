package cont

import "cmp"

// Sum calculates the sum of elements in a slice.
func Sum[S ~[]E, E cmp.Ordered](items S) E {
	var ret E

	for _, item := range items {
		ret += item
	}

	return ret
}
