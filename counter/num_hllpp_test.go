package counter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/counter"
)

// nolint
func FuzzNumHLLPP(f *testing.F) {
	testcases := [...]int{0, 1, 2, 3, 4}
	for _, tc := range testcases {
		f.Add(tc)
	}

	var count uint64

	hll := counter.NewNumHLLPP[int]()

	f.Fuzz(func(t *testing.T, orig int) {
		hll.Add(orig)
		count++
		assert.Equal(t, count, hll.Count())
	})
}

func TestNumHLLPP_Marshal(t *testing.T) {
	t.Parallel()

	hll := counter.NewNumHLLPP[int]()
	hll.Add(3)
	hll.Add(1)
	assert.Equal(t, uint64(2), hll.Count())

	newHll, err := counter.Unmarshal[int](hll.Marshal())
	assert.Nil(t, err)
	assert.Equal(t, uint64(2), newHll.Count())
}

func TestUnMarshal(t *testing.T) {
	t.Parallel()

	_, err := counter.Unmarshal[int]([]byte{})
	assert.NotNil(t, err)
}

func TestNumber2Bytes(t *testing.T) {
	t.Parallel()

	assert.EqualValues(t, []byte{0x1}, counter.Number2Bytes(1))
	assert.EqualValues(t, []byte{0x40, 0x42, 0xf}, counter.Number2Bytes(1000000))
	assert.EqualValues(t, []byte{0x41, 0x42, 0xf}, counter.Number2Bytes(1000001))
	assert.EqualValues(t, []byte{}, counter.Number2Bytes(0))
	assert.EqualValues(t, []byte{0x1f, 0x85, 0xeb, 0x51, 0xb8, 0x1e, 0x9, 0x40}, counter.Number2Bytes(3.14))
}
