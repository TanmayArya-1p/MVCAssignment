package utils

import (
	"sync"
	"sync/atomic"
	"time"
)

type SingleCache[T any] struct {
	Value      T
	IsOccupied atomic.Bool
}

func NewSingleCache[T any]() *SingleCache[T] {
	return &SingleCache[T]{}
}

func (c *SingleCache[T]) Set(value T) {
	c.Value = value
	c.IsOccupied.Store(true)
}

func (c *SingleCache[T]) Delete() {
	c.Value = *new(T)
	c.IsOccupied.Store(false)
}

func (c *SingleCache[T]) Get() *T {
	if !c.IsOccupied.Load() {
		return nil
	}
	return &c.Value
}

type CacheEntry struct {
	Value     any
	ExpiresAt time.Time
	Ticker    *time.Timer
}

type CacheController struct {
	cache    sync.Map
	TTL      time.Duration
	Capacity uint
	count    uint
}

// FOR REFERENCE: customTTL == -1 seconds when we want to use the default TTL defined in the controller struct.
func (c *CacheController) Set(key any, value any, customTTL time.Duration) error {

	if c.count >= c.Capacity {
		return ErrCacheFull
	}

	var ttl time.Duration = c.TTL
	if customTTL != -1*time.Second {
		ttl = customTTL
	}

	prev, loaded := c.cache.Swap(key, CacheEntry{
		Value:     value,
		ExpiresAt: time.Now().Add(c.TTL),
		Ticker: time.AfterFunc(ttl, func() {
			c.Delete(key)
		}),
	})
	if loaded {
		prev.(CacheEntry).Ticker.Stop()
	}
	c.count++
	return nil
}

func (c *CacheController) Get(key any) any {
	val, ok := c.cache.Load(key)
	if !ok {
		return nil
	}
	var cacheVal CacheEntry = val.(CacheEntry)
	if cacheVal.ExpiresAt.Before(time.Now()) {
		c.Delete(key)
		return nil
	}
	return cacheVal.Value
}

func (c *CacheController) Delete(key any) {
	val, loaded := c.cache.LoadAndDelete(key)
	if loaded {
		c.count--
		val.(CacheEntry).Ticker.Stop()
	}
}

func (c *CacheController) Clear() {
	c.cache.Range(func(key any, value any) bool {
		value.(CacheEntry).Ticker.Stop()
		return true
	})
	c.count = 0
	c.cache.Clear()
}

func NewCacheController(ttl time.Duration) *CacheController {
	return &CacheController{
		cache:    sync.Map{},
		TTL:      ttl,
		Capacity: MaxCacheSize,
		count:    0,
	}
}
