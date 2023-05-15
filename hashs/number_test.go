package hashs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/hashs"
)

func TestJSNumber(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0x155dfb1875c8c7, hashs.JSNumber(0x5c955dfb1875c8c7))
	assert.Equal(t, 0x155dfb1875c8c7, hashs.JSNumber(0x5c955dfb1875c8c7))
}

// nolint
func FuzzJSNumber(f *testing.F) {
	f.Add(uint64(5))

	f.Fuzz(func(t *testing.T, number uint64) {
		t.Helper()

		assert.LessOrEqual(t, hashs.JSNumber(number), uint64(1<<53))
	})
}
