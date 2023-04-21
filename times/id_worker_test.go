package times_test

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/times"
)

// nolint: paralleltest
func TestIDWorker_ID(t *testing.T) {
	worker := times.NewIDWorker()

	patches := gomonkey.ApplyFuncReturn(time.Now, lo.Must1(time.Parse("2006-01-01", "2023-01-01")))
	defer patches.Reset()

	ass := assert.New(t)

	ass.Panics(func() {
		worker.ID()
	})
}

func TestIDWorker_IDAndError(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	worker := times.NewIDWorker()
	size := 10000
	ids := make(map[int64]int, size)

	time.Sleep(time.Second)

	for i := 0; i < size; i++ {
		ids[worker.ID()] = i
	}

	ass.Equal(size, len(ids))
}
