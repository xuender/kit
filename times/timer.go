package times

import (
	"context"
	"time"
)

// WithTimer 创建一个协程，其中包含一个计时器和一个回调函数。
// 计时器将在指定的持续时间后触发，并调用回调函数。
// 返回一个停止函数，可以通过调用它来取消计时器和回调函数的执行。
func WithTimer(duration time.Duration, yield func()) func() {
	timer := time.NewTimer(duration)
	stopChannel := make(chan struct{})

	go func() {
		defer timer.Stop()

		select {
		case <-timer.C:
			yield()
		case <-stopChannel:
		}
	}()

	return func() {
		close(stopChannel)
		timer.Stop()
	}
}

// WithContextTimer 是一个函数，它使用一个定时器，并在指定时间后执行给定的yield函数，
// 同时允许通过上下文参数来控制是否取消这个执行。
//
// 参数:
//
//	ctx: 用于取消定时器操作的上下文。
//	duration: 定时器的持续时间，决定何时执行yield函数。
//	yield: 一个函数，当定时器到期时会被调用。
func WithContextTimer(ctx context.Context, duration time.Duration, yield func()) {
	timer := time.NewTimer(duration)

	go func() {
		defer timer.Stop()

		select {
		case <-timer.C:
			yield()
		case <-ctx.Done():
		}
	}()
}
