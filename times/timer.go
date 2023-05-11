package times

import (
	"context"
	"time"
)

func WithTimer(duration time.Duration, yield func()) func() {
	timer := time.NewTimer(duration)
	stopChannel := make(chan struct{})

	go func() {
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

func WithContextTimer(ctx context.Context, duration time.Duration, yield func()) {
	timer := time.NewTimer(duration)

	go func() {
		select {
		case <-timer.C:
			yield()
		case <-ctx.Done():
		}
	}()
}
