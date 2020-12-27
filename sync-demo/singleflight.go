package sync_demo

import (
	"sync"

	"golang.org/x/sync/singleflight"
)

// SingleMap 避免缓存击穿的map
type SingleMap struct {
	// data store
	*sync.Map

	// 防止缓存击穿 avoid cache breakdown
	*singleflight.Group
	getter GetterFunc
}

type GetterFunc func(key string) string

func NewSingleMap(getterFunc GetterFunc) *SingleMap {
	return &SingleMap{
		new(sync.Map),
		new(singleflight.Group),
		getterFunc,
	}

}

func (s *SingleMap) Get(key string) (string, error) {
	//	1. load from map
	load, ok := s.Load(key)
	if ok {
		return load.(string), nil
	}

	// 2. if load fail ,use singleflight to protect getter
	v, err, _ := s.Do(key, func() (interface{}, error) {
		// 2.1 may be load by other goroutines
		load, ok = s.Load(key)
		if ok {
			return load.(string), nil
		}

		// 2.2 load by getter
		value := s.getter(key)
		s.Store(key, value)
		return value, nil
	})

	return v.(string), err
}
