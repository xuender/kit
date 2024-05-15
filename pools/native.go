package pools

import (
	"github.com/xuender/kit/sync"
)

// Native 原生的协程控制.
type Native[T any] struct {
	yield func(T, int)
	group *sync.RoutineGroup
}

// NewNative 新建原生的协程控制.
func NewNative[T any](size uint, yield func(T, int)) *Native[T] {
	return &Native[T]{
		yield,
		sync.NewRoutineGroup(int32(size)),
	}
}

// Post 发送数据.
func (p *Native[T]) Post(elems ...T) {
	for idx, elem := range elems {
		p.group.Incr()
		go p.yield(elem, idx)
	}
}

// Close 关闭协程池.
func (p *Native[T]) Close() {
	p.group.Wait()
}

// Wait 等待完毕.
func (p *Native[T]) Wait() {
	p.group.Wait()
}
