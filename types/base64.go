package types

import (
	"golang.org/x/exp/constraints"
)

// nolint: gochecknoglobals
var (
	_alphabet    = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!_")
	_alphabetMap = make([]int, _max)
)

const (
	_max = 123
)

// nolint: gochecknoinits
func init() {
	for i := 0; i < len(_alphabet); i++ {
		_alphabetMap[_alphabet[i]] = i
	}
}

// NumToB64 数值转换 Base64.
// nolint: gomnd
func NumToB64[N constraints.Integer | constraints.Float](num N) string {
	if num < 0 {
		return "-" + NumToB64(int(num)*-1)
	}

	var (
		n64 = int64(num)
		ret = []byte{}
	)

	for {
		ret = append([]byte{_alphabet[0x3f&n64]}, ret...)

		n64 >>= 6
		if n64 <= 0 {
			break
		}
	}

	return string(ret)
}

// B64ToNum Base64 转换成数值.
func B64ToNum[N constraints.Integer | constraints.Float](str string) N {
	var (
		byt  = []byte(str)
		num  = 0
		sign = 0
	)

	if byt[0] == '-' {
		sign = 1
	}

	for i := sign; i < len(byt); i++ {
		// nolint: gomnd
		num = num*64 + _alphabetMap[str[i]]
	}

	if sign > 0 {
		return N(num * -1)
	}

	return N(num)
}
