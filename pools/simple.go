package pools

type Simple[T any] struct {
	yield func(T, int)
	chans chan T
}

func NewSimple[T any](size int, yield func(T, int)) *Simple[T] {
	ret := &Simple[T]{yield, make(chan T, size)}

	for i := 0; i < size; i++ {
		go ret.run(i)
	}

	return ret
}

func (p *Simple[T]) Post(elems ...T) {
	for _, elem := range elems {
		p.chans <- elem
	}
}

func (p *Simple[T]) run(num int) {
	for elem := range p.chans {
		p.yield(elem, num)
	}
}
