package los

import "github.com/samber/lo"

func EveryNil(elems ...any) bool {
	return lo.EveryBy(elems, func(elem any) bool { return elem == nil })
}

func SomeNil(elems ...any) bool {
	return lo.SomeBy(elems, func(elem any) bool { return elem == nil })
}
