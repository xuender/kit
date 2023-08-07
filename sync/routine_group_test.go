package sync_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/sync"
)

func TestNewRoutineGroup_panic(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, sync.ErrSizeLessZero.Error(), func() {
		sync.NewRoutineGroup(0)
	})
}
