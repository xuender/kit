package times

import (
	"sync"
	"time"
)

func Throttle(yield func(), interval time.Duration) (func(), func()) {
	var (
		mutex   sync.Mutex
		last    time.Time
		timer   *time.Timer
		group   sync.WaitGroup
		firstMu sync.Once
	)

	firstMu.Do(func() {
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
