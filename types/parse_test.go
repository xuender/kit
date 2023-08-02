package types_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/types"
)

func TestParseInteger(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 10, lo.Must1(types.ParseInteger[int]("10")))
}

func TestParseInteger_Float(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 10, lo.Must1(types.ParseInteger[int]("10.3")))
	assert.Equal(t, 11, lo.Must1(types.ParseInteger[int]("10.5")))
}

func TestParseInteger_Error(t *testing.T) {
	t.Parallel()

	_, err := types.ParseInteger[int]("xxfef.3r3r")

	assert.NotNil(t, err)
}

func TestParseFloat(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3.14, lo.Must1(types.ParseFloat[float64]("3.14")))
}

func TestItoa(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "3", types.Itoa(3))
	assert.Equal(t, "3", types.Itoa(3.14))
}

func TestFormatFloat(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "3", types.FormatFloat(3, 3))
	assert.Equal(t, "3.14", types.FormatFloat(3.14, 3))
}

func TestRound(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3, types.Round[int](3.14))
	assert.Equal(t, 3, types.Round[int](2.74))
}

func TestParseIntegerAny(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int]("3")))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int]("3.0")))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int](3)))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int](3.0)))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int](float32(3.0))))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int](uint(3))))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int](int8(3))))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int](uint8(3))))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int](int16(3))))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int](uint16(3))))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int](int32(3))))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int](uint32(3))))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int](int64(3))))
	assert.Equal(t, 3, lo.Must1(types.ParseIntegerAny[int](uint64(3))))
	assert.Equal(t, 102, lo.Must1(types.ParseIntegerAny[int]('f')))
	assert.Equal(t, 1000000, lo.Must1(types.ParseIntegerAny[int]([]byte{0x40, 0x42, 0xf})))

	_, err := types.ParseIntegerAny[int]([]int{})
	assert.NotNil(t, err)
}
