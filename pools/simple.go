package pools

import "sync"

// Simle 简单协程池.
type Simple[T any] struct {
	yield func(T, int)
	queue chan T
	wait  sync.WaitGroup
}

// NewSimle 新建简单协程池.
func NewSimple[T any](size int, yield func(T, int)) *Simple[T] {
	pool := &Simple[T]{yield, make(chan T, size), sync.WaitGroup{}}

	for i := 0; i < size; i++ {
		go pool.run(i)
	}

	return pool
}

// Post 发送数据.
func (p *Simple[T]) Post(elems ...T) {
	p.wait.Add(len(elems))

	for _, elem := range elems {
		p.queue <- elem
	}
}

// Close 关闭协程池.
func (p *Simple[T]) Close() {
	p.wait.Wait()
	close(p.queue)
}

// Wait 等待完毕.
func (p *Simple[T]) Wait() {
	p.wait.Wait()
}

func (p *Simple[T]) run(num int) {
	for elem := range p.queue {
		p.yield(elem, num)
		p.wait.Done()
	}
}
