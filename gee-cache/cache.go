package gee_cache

import (
	"sync"

	"github.com/NothingXiang/go-every-week/gee-cache/lru"
)

type cache struct {
	sync.RWMutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.Lock()
	defer c.Unlock()

	if c.lru == nil {
		c.lru = lru.NewCache(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (ByteView, bool) {
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
