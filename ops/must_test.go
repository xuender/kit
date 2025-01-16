package ops_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/v2/ops"
)

func TestMust(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ops.Must(1, nil)
	ass.Panics(func() {
		ops.Must(1, os.ErrClosed)
	})
	ass.Panics(func() {
		ops.Must(1, os.ErrClosed, "code: %d", 3)
	})
}

func TestMust0(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ops.Must0(nil)
	ass.Panics(func() {
		ops.Must0(os.ErrClosed)
	})
}

func TestMust1(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ops.Must1(1, nil)
	ass.Panics(func() {
		ops.Must1(1, os.ErrClosed)
	})
}

func TestMust2(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ops.Must2(1, "a", nil)
	ass.Panics(func() {
		ops.Must2(1, "a", os.ErrClosed)
	})
}

func TestMust3(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ops.Must3(1, "a", 1, nil)
	ass.Panics(func() {
		ops.Must3(1, "a", 1, os.ErrClosed)
	})
}

func TestMust4(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ops.Must4(true, 1, "a", 1, nil)
	ass.Panics(func() {
		ops.Must4(true, 1, "a", 1, os.ErrClosed)
	})
}

func TestMust5(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ops.Must5(true, 1, "a", 1, 3.3, nil)
	ass.Panics(func() {
		ops.Must5(true, 1, "a", 1, 4, os.ErrClosed)
	})
}

func TestMust6(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ops.Must6('a', true, 1, "a", 1, 3.3, nil)
	ass.Panics(func() {
		ops.Must6('a', true, 1, "a", 1, 4, os.ErrClosed)
	})
}
