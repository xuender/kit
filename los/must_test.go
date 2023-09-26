package los_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/los"
)

func TestMust(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	los.Must(1, nil)
	ass.Panics(func() {
		los.Must(1, os.ErrClosed)
	})
}

func TestMust0(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	los.Must0(nil)
	los.Must0(true)
	ass.Panics(func() {
		los.Must0(os.ErrClosed)
	})
	ass.Panics(func() {
		los.Must0(false)
	})
}

func TestMust1(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	los.Must1(1, nil)
	ass.Panics(func() {
		los.Must1(1, os.ErrClosed)
	})
}

func TestMust2(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	los.Must2(1, "a", nil)
	ass.Panics(func() {
		los.Must2(1, "a", os.ErrClosed)
	})
}

func TestMust3(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	los.Must3(1, "a", 1, nil)
	ass.Panics(func() {
		los.Must3(1, "a", 1, os.ErrClosed)
	})
}

func TestMust4(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	los.Must4(true, 1, "a", 1, nil)
	ass.Panics(func() {
		los.Must4(true, 1, "a", 1, os.ErrClosed)
	})
}

func TestMust5(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	los.Must5(true, 1, "a", 1, 3.3, nil)
	ass.Panics(func() {
		los.Must5(true, 1, "a", 1, 4, 5, os.ErrClosed)
	})
}

func TestMust6(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	los.Must6('a', true, 1, "a", 1, 3.3, nil)
	ass.Panics(func() {
		los.Must6('a', true, 1, "a", 1, 4, 5, os.ErrClosed)
	})
}
