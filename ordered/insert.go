package ordered

import "cmp"

func InsertAes[T cmp.Ordered](slice []T, elem T) []T {
	index := IndexAes(slice, elem)
	right := append([]T{elem}, slice[index:]...)

	return append(slice[:index], right...)
}

func InsertDesc[T cmp.Ordered](slice []T, elem T) []T {
	index := IndexDesc(slice, elem)
	right := append([]T{elem}, slice[index:]...)

	return append(slice[:index], right...)
}
