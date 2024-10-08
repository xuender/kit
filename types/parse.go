package types

import (
	"math"
	"strconv"
	"strings"

	"github.com/xuender/kit/base"
	"golang.org/x/exp/constraints"
)

// ParseInteger 字符串转换成整数.
// nolint
func ParseInteger[T constraints.Integer](str string) (T, error) {
	if strings.ContainsRune(str, '.') {
		float, err := ParseFloat[float64](str)
		if err != nil {
			return 0, err
		}

		return Round[T](float), nil
	}

	i64, err := strconv.ParseInt(str, base.Ten, base.SixtyFour)
	// nolint
	return T(i64), err
}

// ParseIntegerAny 任意类型转换成数值.
// nolint
func ParseIntegerAny[T constraints.Integer](elem any) (T, error) {
	switch num := elem.(type) {
	case string:
		return ParseInteger[T](num)
	case float32:
		return Round[T](num), nil
	case float64:
		return Round[T](num), nil
	case int:
		return T(num), nil
	case uint:
		return T(num), nil
	case int8:
		return T(num), nil
	case uint8:
		return T(num), nil
	case int16:
		return T(num), nil
	case uint16:
		return T(num), nil
	case int32:
		return T(num), nil
	case uint32:
		return T(num), nil
	case int64:
		return T(num), nil
	case uint64:
		return T(num), nil
	case []byte:
		return Bytes2Number[T](num), nil
	}

	return 0, ErrNotNum
}

// ParseFloat 字符串转换成浮点数.
func ParseFloat[T constraints.Float](str string) (T, error) {
	f64, err := strconv.ParseFloat(str, base.SixtyFour)
	// nolint
	return T(f64), err
}

// Itoa 整数转换成字符串.
func Itoa[T constraints.Integer | constraints.Float](num T) string {
	return strconv.Itoa(int(num))
}

func FormatUint[T constraints.Unsigned](num T) string {
	return strconv.FormatUint(uint64(num), base.Ten)
}

// FormatFloat 浮点数格式化成字符串.
func FormatFloat[T constraints.Float | constraints.Integer](num T, prec int) string {
	return strconv.FormatFloat(float64(num), 'g', prec, base.SixtyFour)
}

// Round 四舍五入.
func Round[I constraints.Integer, F constraints.Float](float F) I {
	half := 0.5
	// nolint
	return I(math.Floor(float64(float) + half))
}
