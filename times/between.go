package times

import (
	"context"
	"time"

	"github.com/xuender/kit/base"
)

// Between 时间范围内执行，超出时间终止.
func Between(ctx context.Context, start, stop int, yield func(context.Context)) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if InScope(start, stop) {
				cancelCtx, cancel := context.WithCancel(ctx)

				go yield(cancelCtx)

				time.Sleep(Sleep(stop))
				cancel()
			}

			time.Sleep(time.Minute)
		}
	}
}

// Sleep 计算睡眠时间.
func Sleep(stop int) time.Duration {
	now := time.Now()
	stopTime := time.Date(
		now.Year(), now.Month(), now.Day(),
		stop/base.Hundred, stop%base.Hundred, 0, 0,
		now.Location(),
	)

	if stopTime.Before(now) {
		stopTime = stopTime.AddDate(0, 0, 1)
	}

	return stopTime.Sub(now) + time.Minute
}

// InScope 返回是否在时间范围内.
// nolint: cyclop
func InScope(start, stop int) bool {
	var (
		now    = time.Now()
		hour   = now.Hour()
		minute = now.Minute()
		startH = start / base.Hundred
		stopH  = stop / base.Hundred
		startM = start % base.Hundred
		stopM  = stop % base.Hundred
	)

	if startH > stopH {
		if hour < startH && hour > stopH {
			return false
		}

		if hour == startH && minute < startM {
			return false
		}

		if hour == stopH && minute > stopM {
			return false
		}

		return true
	}

	if hour < startH {
		return false
	}

	if hour == startH && minute < startM {
		return false
	}

	if hour > stopH {
		return false
	}

	if hour == stopH && minute > stopM {
		return false
	}

	return true
}
