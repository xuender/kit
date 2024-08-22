package los

import (
	"slices"
)

func EveryNil(elems ...any) bool {
	return slices.IndexFunc(elems, func(elem any) bool { return elem != nil }) < 0
}

func SomeNil(elems ...any) bool {
	return slices.Index(elems, nil) >= 0
}
