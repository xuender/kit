package sync

import (
	"sync"
	"sync/atomic"
)

// RoutineGroup 协程组，是sync.WaitGroup的增强版本.
type RoutineGroup struct {
	size    int32
	count   atomic.Int32
	lock    sync.Mutex
	wait    sync.WaitGroup
	lockNum atomic.Int32
}

// NewRoutineGroup 协程组，控制协程总数量.
func NewRoutineGroup(size int32) *RoutineGroup {
	if size < 1 {
		panic(ErrGroupLessZero)
	}

	return &RoutineGroup{size: size}
}

// Incr 加1.
func (p *RoutineGroup) Incr() {
	p.wait.Add(1)

	if p.count.Add(1) >= p.size {
		p.lockNum.Add(1)
		p.lock.Lock()
	}
}

// Done 协程完成.
func (p *RoutineGroup) Done() {
	if p.count.Add(-1) < p.size && p.lockNum.Load() >= 0 {
		p.lockNum.Add(-1)
		p.lock.TryLock()
		p.lock.Unlock()
	}

	p.wait.Done()
}

// Wait 等待全部完成.
func (p *RoutineGroup) Wait() {
	p.wait.Wait()
}
