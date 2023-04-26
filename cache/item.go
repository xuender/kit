package cache

import "time"

type Item[V any] struct {
	Data       V
	Expiration int64
}

func (item Item[V]) Expired() bool {
	return item.ExpiredByTime(time.Now().UnixNano())
}

func (item Item[V]) ExpiredByTime(nano int64) bool {
	if item.Expiration == 0 {
		return false
	}

	return nano > item.Expiration
}
