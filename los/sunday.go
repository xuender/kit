package los

// Sunday 算法数组查找.
type Sunday[C comparable] struct {
	shift  map[C]int
	subLen int
	sub    []C
}

// NewSunday 新建 Sunday 算法.
func NewSunday[C comparable](sub []C) *Sunday[C] {
	var (
		shift  = map[C]int{}
		subLen = len(sub)
	)

	for index, value := range sub {
		shift[value] = subLen - index
	}

	return &Sunday[C]{
		shift:  shift,
		subLen: subLen,
		sub:    sub,
	}
}

// IndexOf 查找位置.
func (p Sunday[C]) IndexOf(slice []C) int {
	length := len(slice)

	switch {
	case p.subLen == 0:
		return 0
	case p.subLen > length:
		return -1
	}

	for index := 0; index <= length-p.subLen; {
		for subIndex := 0; p.sub[subIndex] == slice[index+subIndex]; {
			subIndex++

			if subIndex >= p.subLen {
				return index
			}
		}

		if index == length-p.subLen {
			return -1
		}

		if shift, has := p.shift[slice[index+p.subLen]]; has {
			index += shift
		} else {
			index += p.subLen + 1
		}
	}

	return -1
}
