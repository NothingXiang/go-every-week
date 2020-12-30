// package core defines some core function to a cache ,and provides a lru implement
package core

// CacheCore interface defines a a cache what should do
// has some implements ,like FIFO, LFU, LRU, e.g.
type CacheCore interface {
	Get(key string) (Value, bool)
	Add(key string, value Value)
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}
