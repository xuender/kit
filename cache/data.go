package cache

import (
	"sync"
	"time"

	"github.com/xuender/kit/logs"
)

type data[K comparable, V any] struct {
	defaultExpiration time.Duration
	items             map[K]item[V]
	mutex             sync.RWMutex
	done              chan struct{}
}

func (p *data[K, V]) run(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			p.DeleteExpired()
		case <-p.done:
			logs.D.Println("cache finaliz:", &p)

			return
		}
	}
}

func (p *data[K, V]) DeleteExpired() {
	now := time.Now().UnixNano()

	p.mutex.Lock()

	for key, item := range p.items {
		if item.ExpiredByTime(now) {
			delete(p.items, key)
		}
	}

	p.mutex.Unlock()
}
