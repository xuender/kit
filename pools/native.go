package pools

import (
	"context"
	"sync"

	ksync "github.com/xuender/kit/sync"
)

// Native 原生的协程控制.
type Native[I, O any] struct {
	yield func(context.Context, I) O
	group *ksync.RoutineGroup
}

// NewNative 新建原生的协程控制.
func NewNative[I any, O any](size uint, yield func(context.Context, I) O) *Native[I, O] {
	return &Native[I, O]{
		yield,
		ksync.NewRoutineGroup(int32(size)),
	}
}

// Post 发送数据.
func (p *Native[I, O]) Post(ctx context.Context, elems ...I) []O {
	length := len(elems)
	if length == 0 {
		return nil
	}

	ret := make([]O, length)
	wait := &sync.WaitGroup{}

	wait.Add(length)

	for idx, elem := range elems {
		p.group.Incr()

		go p.run(ctx, idx, elem, ret, wait)
	}

	wait.Wait()

	return ret
}

func (p *Native[I, O]) run(ctx context.Context, idx int, elem I, list []O, wait *sync.WaitGroup) {
	list[idx] = p.yield(ctx, elem)

	p.group.Done()
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
