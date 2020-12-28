package gee_cache

import (
	"sync"

	"github.com/NothingXiang/go-every-week/gee-cache/core"
)

// safelyCache is a concurrency safely cache,wrap a core-cache with mutex
type safelyCache struct {
	sync.RWMutex
	lru        core.CacheCore
	cacheBytes int64
}

func (c *safelyCache) add(key string, value ByteView) {
	c.Lock()
	defer c.Unlock()

	// lazy load
	if c.lru == nil {
		c.lru = core.NewLRUCache(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *safelyCache) get(key string) (ByteView, bool) {
	c.RLock()
	defer c.RUnlock()

	if c.lru == nil {
		return ByteView{}, false
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), true
	}

	return ByteView{}, false
}
