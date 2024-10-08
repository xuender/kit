package times_test

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/los"
	"github.com/xuender/kit/times"
)

// nolint: paralleltest
func TestIDWorker_ID(t *testing.T) {
	worker := times.NewIDWorker()

	patches := gomonkey.ApplyFuncReturn(time.Now, los.Must(time.Parse("2006-01-01", "2023-01-01")))
	defer patches.Reset()

	ass := assert.New(t)

	ass.Panics(func() {
		worker.ID()
	})
}

// nolint: paralleltest
func TestIDWorker_IDs(t *testing.T) {
	worker := times.NewIDWorker()

	patches := gomonkey.ApplyFuncReturn(time.Now, los.Must(time.Parse("2006-01-01", "2023-01-01")))
	defer patches.Reset()

	ass := assert.New(t)

	ass.Panics(func() {
		worker.IDs(10)
	})
}

func TestIDWorker_IDAndError(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	worker := times.NewIDWorker()
	size := 10000
	ids := make(map[int64]int, size)

	time.Sleep(time.Second)

	for i := range size {
		ids[worker.ID()] = i
	}

	ass.Len(ids, size)
}

func TestIDWorker_IDsAndError(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	worker := times.NewIDWorker()
	size := 10000
	ids := make(map[int64]int, size)

	time.Sleep(time.Second)

	for i, uid := range worker.IDs(size) {
		ids[uid] = i
	}

	ass.Len(ids, size)
}
