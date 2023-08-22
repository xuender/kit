package ordered

import (
	"golang.org/x/exp/constraints"
)

type Set[T constraints.Ordered] []T

func NewSet[T constraints.Ordered](slice ...T) Set[T] {
	if len(slice) <= 1 {
		return Set[T](slice)
	}

	set := Set[T]{}

	for _, elem := range slice {
		set.Add(elem)
	}

	return set
}

func (p Set[T]) Has(elem T) bool {
	length := len(p)

	switch {
	case length == 0, elem < p[0], elem > p[length-1]:
		return false
	case length == 1:
		return p[0] == elem
	}

	half := length / _two
	if p[half:].Has(elem) {
		return true
	}

	return p[:half].Has(elem)
}

func (p *Set[T]) Add(elem T) {
	index := IndexSet(*p, elem)
	if index < 0 {
		return
	}

	right := append([]T{elem}, (*p)[index:]...)
	*p = append((*p)[:index], right...)
}
