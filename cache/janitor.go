package cache

import (
	"time"

	"github.com/xuender/kit/logs"
)

type janitor struct {
	interval time.Duration
	stop     chan bool
}

func (p *janitor) Run(del deleter) {
	ticker := time.NewTicker(p.interval)

	for {
		select {
		case <-ticker.C:
			del.DeleteExpired()
		case <-p.stop:
			ticker.Stop()
			logs.W.Println(&p)

			return
		}
	}
}

func runJanitor[K comparable, V any](cache *Cache[K, V], interval time.Duration) {
	jan := &janitor{
		interval: interval,
		stop:     make(chan bool),
	}

	cache.janitor = jan
	go jan.Run(cache)
}
