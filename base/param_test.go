package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/base"
)

func TestParam1(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, base.Param1(1, 2))
}

func TestParam2(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 2, base.Param2(1, 2))
}

func TestParam3(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 3, base.Param3(1, 2, 3))
}
