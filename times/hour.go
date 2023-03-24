package times

import (
	"time"
)

// Hour 整点运行，返回取消方法.
func Hour(yield func()) func() bool {
	now := time.Now()
	next := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		0, 0, 0,
		time.Local,
	)

	for next.Before(time.Now()) || next.Equal(time.Now()) {
		next = next.Add(time.Hour)
	}

	timer := time.AfterFunc(next.Sub(now), func() {
		yield()
	})

	return timer.Stop
}
