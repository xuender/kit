package cache

import (
	"runtime"
	"sync"
	"time"
)

const (
	// DefaultExpiration 默认过期时间.
	DefaultExpiration time.Duration = 0
	// NoExpiration 不过期.
	NoExpiration time.Duration = -1
)

// Cache 缓存.
type Cache[K comparable, V any] struct{ *data[K, V] }

// New 新建缓存, 设置默认过期时间和过期检查周期.
func New[K comparable, V any](defaultExpiration, interval time.Duration) *Cache[K, V] {
	cacheData := &data[K, V]{
		defaultExpiration: defaultExpiration,
		items:             make(map[K]item[V]),
		mutex:             sync.RWMutex{},
		done:              make(chan struct{}),
	}
	newCache := &Cache[K, V]{cacheData}

	if interval > 0 {
		go cacheData.run(interval)

		runtime.SetFinalizer(newCache, stop[K, V])
	}

	return newCache
}

func stop[K comparable, V any](cache *Cache[K, V]) {
	cache.done <- struct{}{}
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

// Get 获取元素.
func (p *Cache[K, V]) Get(key K) (V, bool) {
	var value V

	p.mutex.RLock()
	item, found := p.items[key]
	p.mutex.RUnlock()

	if !found {
		return value, false
	}

	if item.Expired() {
		return value, false
	}

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
