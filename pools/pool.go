package pools

import (
	"runtime"

	"github.com/xuender/kit/logs"
)

// Pool Goroutine 池.
type Pool[I, O any] struct{ *data[I, O] }

// New 新建 Goroutine 池.
func New[I, O any](size int, yield func(I, int) O) *Pool[I, O] {
	poolData := &data[I, O]{
		chans: make(chan *job[I, O], size),
		yield: yield,
	}
	pool := &Pool[I, O]{poolData}

	for num := 0; num < size; num++ {
		go poolData.run(num)
	}

	runtime.SetFinalizer(pool, stop[I, O])

	return pool
}

func stop[I, O any](pool *Pool[I, O]) {
	close(pool.chans)
	logs.D.Println("pool finaliz:", &pool)
}
