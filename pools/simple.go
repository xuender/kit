package pools

import (
	"runtime"

	"github.com/xuender/kit/logs"
)

type Simple[T any] struct{ *simpleData[T] }

func NewSimple[T any](size int, yield func(T, int)) *Simple[T] {
	data := &simpleData[T]{yield, make(chan T, size)}
	pool := &Simple[T]{data}

	for i := 0; i < size; i++ {
		go data.run(i)
	}

	runtime.SetFinalizer(pool, simpleStop[T])

	return pool
}

type simpleData[T any] struct {
	yield func(T, int)
	chans chan T
}

func simpleStop[T any](simple *Simple[T]) {
	close(simple.chans)

	logs.D.Println("simple stop")
}

func (p *simpleData[T]) Post(elems ...T) {
	for _, elem := range elems {
		p.chans <- elem
	}
}

func (p *simpleData[T]) run(num int) {
	for elem := range p.chans {
		p.yield(elem, num)
	}
}
