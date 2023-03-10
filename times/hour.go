package times

import (
	"time"

	"github.com/xuender/kit/base"
)

// Hour 整点运行，返回取消方法.
func Hour(yield func()) func() bool {
	now := time.Now()
	unix := now.Add(time.Hour).Unix()
	unix -= int64(now.Second())
	unix -= int64(base.Sixty * now.Minute())
	next := time.Unix(unix, 0)

	timer := time.AfterFunc(next.Sub(now), func() {
		yield()
		Hour(yield)
	})

	return timer.Stop
}
