package cache

import "time"

type item[V any] struct {
	value      V
	expiration int64
}

func (p item[V]) Expired() bool {
	return p.ExpiredByTime(time.Now().UnixNano())
}

func (p item[V]) ExpiredByTime(nano int64) bool {
	if p.expiration == 0 {
		return false
	}

	return nano > p.expiration
}
