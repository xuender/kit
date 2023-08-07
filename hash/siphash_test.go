package hash_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/hash"
)

func TestSipHash128(t *testing.T) {
	t.Parallel()

	h1, h2 := hash.SipHash128([]byte("123"))

	assert.Equal(t, uint64(0x78a608799f2d09cf), h1)
	assert.Equal(t, uint64(0xa13d04dbacc077f8), h2)
}

func TestSipHash64(t *testing.T) {
	t.Parallel()

	h := hash.SipHash64([]byte("123"))

	assert.Equal(t, uint64(0x822983866c7d3daf), h)
}

func FuzzSipHashNumber(f *testing.F) {
	f.Add([]byte("test"))

	f.Fuzz(func(t *testing.T, data []byte) {
		t.Helper()

		hash := hash.SipHashNumber(data)

		assert.LessOrEqual(t, hash, uint64(1<<53))
		assert.GreaterOrEqual(t, hash, uint64(100))
	})
}

func TestSipHash32(t *testing.T) {
	t.Parallel()

	h := hash.SipHash32([]byte("123"))

	assert.Equal(t, uint32(0x82298386), h)
}
