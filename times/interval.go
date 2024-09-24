package times

import "time"

// Interval 返回一个函数，该函数确保每次调用时都经过一段特定的时间间隔。
// interval 参数指定了两次调用之间所需的最小时间间隔。
func Interval(interval time.Duration) func() {
	before := time.Now()

	return func() {
		if elapsed := time.Since(before); interval > elapsed {
			time.Sleep(interval - elapsed)
		}

		before = time.Now()
	}
}
