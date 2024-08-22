package los

import "slices"

func IndexAll[S ~[]E, E comparable](slice S, elem E) []int {
	result := []int{}

	for start := 0; ; start++ {
		idx := slices.Index(slice[start:], elem)
		if idx < 0 {
			return result
		}

		start += idx
		result = append(result, start)
	}
}

func IndexAllFunc[S ~[]E, E any](slice S, predicate func(E) bool) []int {
	result := []int{}

	for start := 0; ; start++ {
		idx := slices.IndexFunc(slice[start:], predicate)
		if idx < 0 {
			return result
		}

		start += idx
		result = append(result, start)
	}
}
