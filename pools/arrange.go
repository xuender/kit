package pools

import (
	"sort"

	"github.com/xuender/kit/los"
	"golang.org/x/exp/constraints"
)

// Arrange 整理成员执行顺序，key相等的避免同时执行.
func Arrange[T any, O constraints.Ordered](elems []T, getKey func(T) O) []T {
	groups := map[O][]T{}

	for _, elem := range elems {
		key := getKey(elem)

		if group, has := groups[key]; has {
			groups[key] = append(group, elem)
		} else {
			groups[key] = []T{elem}
		}
	}

	keys := make([]O, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		leni := len(groups[keys[i]])
		lenj := len(groups[keys[j]])

		if leni == lenj {
			return los.SampleBool()
		}

		return lenj < leni
	})

	ret := make([]T, len(elems))
	index := 0

	for i := 0; i < len(elems); i++ {
		for _, key := range keys {
			if len(groups[key]) > i {
				ret[index] = groups[key][i]
				index++
			}
		}
	}

	return ret
}
