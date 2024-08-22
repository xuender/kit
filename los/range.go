package los

import (
	"golang.org/x/exp/constraints"
)

func Range[T constraints.Integer](elementNum uint) []T {
	result := make([]T, elementNum)
	for num := range elementNum {
		result[num] = T(num)
	}

	return result
}
