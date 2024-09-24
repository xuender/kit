package times

import (
	"sync"
	"time"
)

// Throttle 创建一个节流控制器，用于控制函数的调用频率。
//
// 参数:
//   - yield: 要节流的函数。
//   - interval: 间隔时间。
//
// 返回值:
//   - 第一个返回的函数用于触发节流事件。
//   - 第二个返回的函数用于阻塞调用，直到节流函数执行完毕。
func Throttle(yield func(), interval time.Duration) (func(), func()) {
	var (
		mutex sync.Mutex
		last  time.Time
		timer *time.Timer
		group sync.WaitGroup
		first sync.Once
	)

	first.Do(func() {
		last = time.Now()
	})

	call := func() {
		yield()

		last = time.Now()
		timer = nil

		group.Done()
	}

	return func() {
			mutex.Lock()
			defer mutex.Unlock()

			group.Add(1)

			if elapsed := time.Since(last); elapsed < interval {
				if timer == nil {
					timer = time.AfterFunc(interval-elapsed, func() {
						mutex.Lock()
						defer mutex.Unlock()

						call()
					})
				} else {
					group.Done()
				}

				return
			}

			if timer != nil {
				timer.Stop()
			}

			call()
		}, func() {
			group.Wait()
		}
}
