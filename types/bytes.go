package types

import (
	"encoding/binary"
	"math"

	"github.com/xuender/kit/base"
	"golang.org/x/exp/constraints" // nolint
)

func Number2Bytes[T constraints.Integer | constraints.Float](num T) []byte {
	switch (any)(num).(type) {
	case float32, float64:
		return Number2Bytes(math.Float64bits(float64(num)))
	}

	data := make([]byte, base.Eight)
	binary.LittleEndian.PutUint64(data, uint64(num))

	for i := base.Eight; i > 0; i-- {
		if data[i-1] > 0 {
			return data[:i]
		}
	}

	return nil
}

func Bytes2Number[T constraints.Integer | constraints.Float](data []byte) T {
	bytes := make([]byte, base.Eight)
	copy(bytes, data)
	num := binary.LittleEndian.Uint64(bytes)

	tt := new(T)

	switch (any)(*tt).(type) {
	case float32, float64:
		return T(math.Float64frombits(num))
	}

	return T(num)
}
