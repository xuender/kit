package cont

import (
	"github.com/xuender/kit/v2/types"
)

// Mean calculates the mean of the elements in the slice.
func Mean[S ~[]E, E types.Number](items S) E {
	if len(items) == 0 {
		return 0
	}

	return Sum(items) / E(len(items))
}
