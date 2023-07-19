package pools

// Simle 简单协程池.
type Simple[T any] struct {
	yield func(T, int)
	queue chan T
}

// NewSimle 新建简单协程池.
func NewSimple[T any](size int, yield func(T, int)) *Simple[T] {
	pool := &Simple[T]{yield, make(chan T, size)}

	for i := 0; i < size; i++ {
		go pool.run(i)
	}

	return pool
}

// Post 发送数据.
func (p *Simple[T]) Post(elems ...T) {
	for _, elem := range elems {
		p.queue <- elem
	}
}

// Close 关闭协程池.
func (p *Simple[T]) Close() {
	close(p.queue)
}

func (p *Simple[T]) run(num int) {
	for elem := range p.queue {
		p.yield(elem, num)
	}
}
