package cache

import (
	"sync"
	"time"

	"github.com/xuender/kit/logs"
)

const (
	// DefaultExpiration 默认过期时间.
	DefaultExpiration time.Duration = 0
	// NoExpiration 不过期.
	NoExpiration time.Duration = -1
)

// Cache 缓存.
type Cache[K comparable, V any] struct {
	defaultExpiration time.Duration
	items             map[K]item[V]
	mutex             sync.RWMutex
	done              chan struct{}
}

// New 新建缓存, 设置默认过期时间和过期检查周期.
func New[K comparable, V any](defaultExpiration, interval time.Duration) *Cache[K, V] {
	ret := &Cache[K, V]{
		defaultExpiration,
		make(map[K]item[V]),
		sync.RWMutex{},
		make(chan struct{}),
	}

	if interval > 0 {
		go ret.run(interval)
	}

	return ret
}

func (p *Cache[K, V]) Close() {
	close(p.done)
}

// NewStringKey 新建字符串键值的缓存.
func NewStringKey[V any](defaultExpiration, interval time.Duration) *Cache[string, V] {
	return New[string, V](defaultExpiration, interval)
}

// Set 设置元素.
func (p *Cache[K, V]) Set(key K, value V) {
	p.SetDuration(key, value, DefaultExpiration)
}

// SetDuration 设置元素及过期时间.
func (p *Cache[K, V]) SetDuration(key K, value V, expiration time.Duration) {
	var exp int64

	if expiration == DefaultExpiration {
		expiration = p.defaultExpiration
	}

	if expiration > 0 {
		exp = time.Now().Add(expiration).UnixNano()
	}

	p.mutex.Lock()
	p.items[key] = item[V]{
		value:      value,
		expiration: exp,
	}
	p.mutex.Unlock()
}

// GetNoExtension 读元素不延期.
func (p *Cache[K, V]) GetNoExtension(key K) (V, bool) {
	var zero V

	p.mutex.RLock()
	item, has := p.items[key]
	p.mutex.RUnlock()

	if !has {
		return zero, false
	}

	if item.Expired() {
		return zero, false
	}

	return item.value, true
}

func (p *Cache[K, V]) Has(key K) bool {
	if len(p.items) == 0 {
		return false
	}

	p.mutex.RLock()
	item, has := p.items[key]
	p.mutex.RUnlock()

	return has && item.Expired()
}

// Get 获取元素并刷新.
func (p *Cache[K, V]) Get(key K) (V, bool) {
	var zero V

	p.mutex.RLock()
	item, has := p.items[key]
	p.mutex.RUnlock()

	if !has {
		return zero, false
	}

	if item.Expired() {
		return zero, false
	}

	p.mutex.Lock()
	item.expiration = time.Now().Add(p.defaultExpiration).UnixNano()
	p.items[key] = item
	p.mutex.Unlock()

	return item.value, true
}

func (p *Cache[K, V]) Delete(key K) {
	p.mutex.Lock()
	delete(p.items, key)
	p.mutex.Unlock()
}

// Iterate 迭代.
func (p *Cache[K, V]) Iterate(yield func(K, V) error) error {
	now := time.Now().UnixNano()

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for key, item := range p.items {
		if item.ExpiredByTime(now) {
			continue
		}

		if err := yield(key, item.value); err != nil {
			return err
		}
	}

	return nil
}

// Len 元素数量.
func (p *Cache[K, V]) Len() int {
	p.DeleteExpired()
	p.mutex.RLock()
	count := len(p.items)
	p.mutex.RUnlock()

	return count
}

func (p *Cache[K, V]) run(interval time.Duration) {
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

func (p *Cache[K, V]) DeleteExpired() {
	now := time.Now().UnixNano()

	p.mutex.Lock()

	for key, item := range p.items {
		if item.ExpiredByTime(now) {
			delete(p.items, key)
		}
	}

	p.mutex.Unlock()
}
