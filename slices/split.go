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
	ret := make([][]T, 0, len(collection)/2+1)
	start := 0

	for end, elem := range collection {
		if isSeparator(elem, end) {
			ret = append(ret, collection[start:end])
			start = end + 1
		}
	}

	ret = append(ret, collection[start:])

	return ret
}
