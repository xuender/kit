package times

import (
	"sync"
	"time"
)

// Debounce 生成一个防抖动函数，用于在一段时间内避免重复触发事件。
// 参数:
//   - yield: 要防抖动的函数。
//   - wait: 等待时间，用于确定在触发事件后是否执行函数。
//
// 返回值:
//   - 第一个返回的函数用于触发防抖动事件。
//   - 第二个返回的函数用于阻塞调用，直到防抖动函数执行完毕。
func Debounce(yield func(), wait time.Duration) (func(), func()) {
	var (
		timer *time.Timer
		mutex sync.Mutex
	)

	cond := sync.NewCond(&mutex)
	call := func() {
		mutex.Lock()
		defer mutex.Unlock()

		yield()
		cond.Broadcast()
	}
	trigger := func() {
		mutex.Lock()
		defer mutex.Unlock()

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(wait, call)
	}
	waitFunc := func() {
		mutex.Lock()
		defer mutex.Unlock()

		if timer != nil {
			cond.Wait()
		}
	}

	return trigger, waitFunc
}
