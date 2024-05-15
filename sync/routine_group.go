package sync

import (
	"sync"
)

// nolint
var _none = struct{}{}

// RoutineGroup 协程组，是sync.WaitGroup的增强版本.
type RoutineGroup struct {
	ch chan struct{}
	wg sync.WaitGroup
}

// NewRoutineGroup 协程组，控制协程总数量.
func NewRoutineGroup(size int32) *RoutineGroup {
	if size < 1 {
		panic(ErrSizeLessZero)
	}

	return &RoutineGroup{
		ch: make(chan struct{}, size),
		wg: sync.WaitGroup{},
	}
}

// Incr 加1.
func (p *RoutineGroup) Incr() {
	p.wg.Add(1)
	p.ch <- _none
}

// Done 协程完成.
func (p *RoutineGroup) Done() {
	<-p.ch
	p.wg.Done()
}

// Wait 等待全部完成.
func (p *RoutineGroup) Wait() {
	p.wg.Wait()
}
