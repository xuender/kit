package cont

import "github.com/xuender/kit/v2/types"

func Find[S ~[]E, E any](slice S, checker types.Checker[E]) (E, bool) {
	for _, item := range slice {
		if checker(item) {
			return item, true
		}
	}

	var zero E

	return zero, false
}
