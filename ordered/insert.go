package ordered

import "golang.org/x/exp/constraints"

func InsertAes[T constraints.Ordered](slice []T, elem T) []T {
	index := IndexAes(slice, elem)
	right := append([]T{elem}, slice[index:]...)

	return append(slice[:index], right...)
}

func InsertDesc[T constraints.Ordered](slice []T, elem T) []T {
	index := IndexDesc(slice, elem)
	right := append([]T{elem}, slice[index:]...)

	return append(slice[:index], right...)
}
