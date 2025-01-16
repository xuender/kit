package cont

import "github.com/xuender/kit/v2/types"

// Reduce applies a reducer function to all elements of a slice and returns a single value.
// It iterates over the slice, applying the reducer to accumulate a result.
// The first element initializes the accumulator.
func Reduce[S ~[]E, E any](items S, reducer types.Reducer[E]) E {
	var ret E

	for idx, item := range items {
		if idx == 0 {
			ret = item
		} else {
			ret = reducer(ret, item)
		}
	}

	return ret
}
