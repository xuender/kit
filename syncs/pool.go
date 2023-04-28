package syncs

import (
	"runtime"

	"github.com/xuender/kit/logs"
)

// Pool Goroutine 池.
type Pool[I, O any] struct{ *data[I, O] }

// NewPool 新建 Goroutine 池.
func NewPool[I, O any](size int, yield func(I, int) O) *Pool[I, O] {
	pool := &data[I, O]{
		input: make(chan *job[I, O], size),
		yield: yield,
	}
	ret := &Pool[I, O]{pool}

	for i := 0; i < size; i++ {
		go pool.run(i)
	}

	runtime.SetFinalizer(ret, stop[I, O])

	return ret
}

func stop[I, O any](pool *Pool[I, O]) {
	logs.D.Println("pool finaliz:", &pool)
	close(pool.input)
}
