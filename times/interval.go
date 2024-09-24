package times

import "time"

func Interval(interval time.Duration) func() {
	before := time.Now()

	return func() {
		if elapsed := time.Since(before); interval > elapsed {
			time.Sleep(interval - elapsed)
		}

		before = time.Now()
	}
}
