package cont

// Fill assigns the value `val` to elements of the slice from index `start` to `end`.
// If `start` and `end` are not provided, it fills the entire slice.
// If only `start` is provided, it fills from `start` to the end of the slice.
func Fill[S ~[]E, E any](slice S, val E, startAndEnd ...int) {
	end := len(slice)
	if end == 0 {
		return
	}

	start := 0

	switch len(startAndEnd) {
	case 0:
	case 1:
		start = startAndEnd[0]
	default:
		start = startAndEnd[0]
		end = startAndEnd[1]
	}

	for idx := start; idx < end; idx++ {
		slice[idx] = val
	}
}
