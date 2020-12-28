package gee_cache

import (
	"errors"
	"log"
	"sync"
)

var (
	// mu protect global groups instance
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

var (
	KeyRequiredErr = errors.New("key is required")
)

// Group is a cache namespace and associated data loaded spread over
type Group struct {
	name      string
	getter    Getter
	mainCache safelyCache
}

// NewGroup create a new instance of Group
// name: cache namespace
// cacheBytes: max cache size
// getter : a callback func when key is not exists it will be called ,could not nil
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil getter")
	}
	mu.Lock()
	defer mu.Unlock()

	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: safelyCache{cacheBytes: cacheBytes},
	}
	groups[name] = g

	return g
}

// GetGroup returns a group by name,return nil if group not exists
func GetGroup(name string) *Group {
	mu.RLock()
	defer mu.RUnlock()

	return groups[name]
}

func (g *Group) Get(key string) (res ByteView, err error) {
	if key == "" {
		return res, KeyRequiredErr
	}
	v, ok := g.mainCache.get(key)
	if ok {
		log.Println("[gee-safelyCache] hit")
		return v, nil
	}

	return g.load(key)

}

// load 预留位置，当从本地获取失败的时候，可以从远程节点获取
func (g *Group) load(key string) (ByteView, error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	get, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{b: cloneBytes(get)}

	g.populateCache(key, value)

	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
