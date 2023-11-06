package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/types"
)

func TestNumber2Bytes(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []byte{0x1}, types.Number2Bytes(1))
	assert.Equal(t, []byte{0x40, 0x42, 0xf}, types.Number2Bytes(1000000))
	assert.Equal(t, []byte{0x41, 0x42, 0xf}, types.Number2Bytes(1000001))
	assert.Equal(t, []byte(nil), types.Number2Bytes(0))
	assert.Equal(t, []byte{0x1f, 0x85, 0xeb, 0x51, 0xb8, 0x1e, 0x9, 0x40}, types.Number2Bytes(3.14))
}

// nolint
func FuzzNumber2Bytes(f *testing.F) {
	testcases := []int{0, 1, 2, 3, 4}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig int) {
		rev := types.Number2Bytes(orig)
		assert.Equal(t, orig, types.Bytes2Number[int](rev))
	})
}

func TestBytes2Number(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, types.Bytes2Number[int]([]byte{0x1}))
	assert.Equal(t, 0, types.Bytes2Number[int]([]byte{}))
	assert.InEpsilon(t, 3.14, types.Bytes2Number[float64]([]byte{0x1f, 0x85, 0xeb, 0x51, 0xb8, 0x1e, 0x9, 0x40}), 2)
}
