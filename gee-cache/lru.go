package gee_cache

import (
	"container/list"
)

// Cache is a LRU cache. It is not safe for concurrent access.
type Cache struct {
	maxBytes uint
	nBytes   uint
	ll       *list.List
	cache    map[string]*list.Element
	// optional and executed when an entry is purged.
	OnEvicted func(key string, value interface{})
}

func NewCache(maxBytes uint, onEvicted func(key string, value interface{})) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() uint
}

// Get look ups a key's value
func (c *Cache) Get(key string) (v Value, ok bool) {
	ele, ok := c.cache[key]
	if ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return nil, false
}

// RemoveOldest removes the oldest item
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)

		c.nBytes -= uint(len(kv.key)) + kv.value.Len()

		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add adds a value to the cache.
func (c *Cache) Add(key string, value Value) {
	// 如果是原来已经存在的key,更新对应的entry;不存在的话则插入
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)

		c.nBytes += value.Len() - kv.value.Len()
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key: key, value: value})
		c.cache[key] = ele
		c.nBytes += uint(len(key)) + value.Len()
	}

	// 淘汰
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
}

// Len the number of cache entries
func (c *Cache) Len() int {
	return c.ll.Len()
}
