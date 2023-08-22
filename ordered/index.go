package ordered

import "golang.org/x/exp/constraints"

const _two = 2

func IndexAes[T constraints.Ordered](slice []T, elem T) int {
	length := len(slice)

	switch {
	case length == 0, elem <= slice[0]:
		return 0
	case length == 1:
		return 1
	case elem >= slice[length-1]:
		return length
	}

	half := length / _two
	if i := IndexAes(slice[half:], elem); i > 0 {
		return half + i
	}

	return IndexAes(slice[:half], elem)
}

func IndexSet[T constraints.Ordered](slice []T, elem T) int {
	length := len(slice)

	switch {
	case length == 0, elem < slice[0]:
		return 0
	case elem == slice[0], elem == slice[length-1]:
		return -1
	case length == 1:
		return 1
	case elem > slice[length-1]:
		return length
	}

	half := length / _two
	index := IndexSet(slice[half:], elem)

	switch {
	case index < 0:
		return index
	case index > 0:
		return half + index
	default:
		return IndexSet(slice[:half], elem)
	}
}

func IndexDesc[T constraints.Ordered](slice []T, elem T) int {
	length := len(slice)

	switch {
	case length == 0, elem >= slice[0]:
		return 0
	case length == 1:
		return 1
	case elem <= slice[length-1]:
		return length
	}

	half := length / _two
	if i := IndexDesc(slice[half:], elem); i > 0 {
		return half + i
	}

	return IndexDesc(slice[:half], elem)
}
