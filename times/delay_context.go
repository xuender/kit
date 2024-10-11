package times

import (
	"context"
	"sync"
	"time"
)

// DelayContext 返回一个函数，该函数用于在给定间隔时间后处理上下文。
// 这个函数的主要目的是在满足指定间隔时间之前延迟上下文的处理。
func DelayContext(interval time.Duration) func(context.Context) error {
	var (
		mutex sync.Mutex
		timer *time.Timer
	)

	last := time.Now()

	return func(ctx context.Context) error {
		mutex.Lock()
		defer mutex.Unlock()

		if elapsed := time.Since(last); interval > elapsed {
			if timer == nil {
				timer = time.NewTimer(interval - elapsed)
			} else {
				timer.Reset(interval - elapsed)
			}

			select {
			case <-ctx.Done():
				timer.Stop()
				timer = nil

				return ctx.Err()
			case <-timer.C:
			}
		}

		last = time.Now()

		return nil
	}
}
