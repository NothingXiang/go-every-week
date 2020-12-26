package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash Hash
	// 虚拟节点倍数
	replicas int
	// 哈希环
	keys []int
	// 虚拟节点hash值与真实节点名称的映射
	hashMap map[int]string
}

func NewMap(replicas int, hash Hash) *Map {
	if hash == nil {
		// default testHashMap method
		hash = crc32.ChecksumIEEE
	}
	return &Map{
		hash:     hash,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
}

// 添加真实节点到map中，每个key对应一个真实节点
// Add adds some keys yo the testHashMap
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			// 通过编号区分不同虚拟节点
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			// 将虚拟节点添加到环
			m.keys = append(m.keys, hash)
			// 维护虚拟节点到真实节点的映射
			m.hashMap[hash] = key
		}
	}
	// 排序,递增
	sort.Ints(m.keys)
}

// Get gets the closest item in the testHashMap to the provided key
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// binary search
	// idx是满足func的最小index ,如果没有满足的，返回n
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	// keys 其实是一个环形结构，如果找了一轮都找不到，那就返回keys[0]
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
