package slices

import (
	"github.com/samber/lo"
)

// SplitStr 字符串根据分隔符分解.
func SplitStr(str string, separators ...rune) []string {
	splits := Split([]rune(str), func(elem rune, _ int) bool { return lo.Contains(separators, elem) })
	ret := make([]string, len(splits))

	for index, spl := range splits {
		ret[index] = string(spl)
	}

	return ret
}

// Split 切片分解.
func Split[T comparable](collection []T, isSeparator func(T, int) bool) [][]T {
	ret := [][]T{}
	tmp := []T{}

	for index, elem := range collection {
		if isSeparator(elem, index) {
			ret = append(ret, tmp)
			tmp = []T{}

			continue
		}

		tmp = append(tmp, elem)
	}

	ret = append(ret, tmp)

	return ret
}
