package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/base"
)

func TestOnly21(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, base.Only21(1, 2))
}

func TestOnly22(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 2, base.Only22(1, 2))
}

func TestOnly31(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, base.Only31(1, 2, 3))
}

func TestOnly32(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 2, base.Only32(1, 2, 3))
}

func TestOnly33(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3, base.Only33(1, 2, 3))
}
