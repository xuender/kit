package pools

import (
	"sync"

	ksync "github.com/xuender/kit/sync"
)

// Native 原生的协程控制.
type Native[I, O any] struct {
	yield func(I, int) O
	group *ksync.RoutineGroup
}

// NewNative 新建原生的协程控制.
func NewNative[I any, O any](size uint, yield func(I, int) O) *Native[I, O] {
	return &Native[I, O]{
		yield,
		ksync.NewRoutineGroup(int32(size)),
	}
}

// Post 发送数据.
func (p *Native[I, O]) Post(elems ...I) []O {
	length := len(elems)
	if length == 0 {
		return nil
	}

	wait := &sync.WaitGroup{}

	wait.Add(length)

	ret := make([]O, length)

	for idx, elem := range elems {
		p.group.Incr()

		go p.run(idx, elem, ret, wait)
	}

	wait.Wait()

	return ret
}

func (p *Native[I, O]) run(idx int, elem I, list []O, wait *sync.WaitGroup) {
	list[idx] = p.yield(elem, idx)

	wait.Done()
}

// Close 关闭协程池.
func (p *Native[I, O]) Close() {
	p.group.Wait()
}

// Wait 等待完毕.
func (p *Native[I, O]) Wait() {
	p.group.Wait()
}
