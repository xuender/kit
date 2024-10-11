package times

import (
	"sync"
	"time"
)

// Delay 返回一个闭包，该闭包在调用时确保每次执行之间至少间隔指定的 duration。
// 这对于需要定期执行任务或限制操作频率的场景非常有用。
func Delay(interval time.Duration) func() {
	var mutex sync.Mutex

	last := time.Now()

	return func() {
		mutex.Lock()
		defer mutex.Unlock()

		if elapsed := time.Since(last); elapsed < interval {
			time.Sleep(interval - elapsed)
		}

		last = time.Now()
	}
}
