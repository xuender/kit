package sync

import (
	"sync"
	"sync/atomic"
)

// RoutineGroup 协程组，是sync.WaitGroup的增强版本.
type RoutineGroup struct {
	group int32
	count atomic.Int32
	lock  sync.Mutex
	wait  sync.WaitGroup
}

// NewRoutineGroup 协程组，控制协程总数量.
func NewRoutineGroup(group int32) *RoutineGroup {
	if group < 1 {
		panic(ErrGroupLessZero)
	}

	return &RoutineGroup{
		group: group,
		count: atomic.Int32{},
		lock:  sync.Mutex{},
		wait:  sync.WaitGroup{},
	}
}

// Incr 加1.
func (p *RoutineGroup) Incr() {
	p.wait.Add(1)

	if p.count.Add(1) >= p.group {
		p.lock.Lock()
	}
}

// Done 协程完成.
func (p *RoutineGroup) Done() {
	if p.count.Add(-1) < p.group {
		p.lock.TryLock()
		p.lock.Unlock()
	}

	p.wait.Done()
}

// Wait 等待全部完成.
func (p *RoutineGroup) Wait() {
	p.wait.Wait()
}
