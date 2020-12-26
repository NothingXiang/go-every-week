package consistenthash

import (
	"strconv"
	"testing"
)

var (
	// 用于测试用的简单哈希函数，数字值会直接转为对应的哈希值，非数字值转为0
	testHash = func(data []byte) uint32 {
		i, _ := strconv.Atoi(string(data))
		return uint32(i)
	}
)

func TestMap_Get(t *testing.T) {
	testHashMap := NewMap(3, testHash)
	// 添加虚拟节点 2, 4, 6, 8, 12, 14, 16, 18, 22, 24, 26, 28
	testHashMap.Add("2", "4", "6", "8")

	tests := []struct {
		name   string
		fields *Map
		key    string
		want   string
	}{
		{
			"hash-equal-node",
			testHashMap,
			"2",
			"2",
		},
		{
			"hash-equal-node",
			testHashMap,
			"14",
			"4",
		},
		{
			"hash-equal-node",
			testHashMap,
			"26",
			"6",
		},

		{
			"hash-nearest-node",
			testHashMap,
			"11",
			"2",
		},
		{
			"hash-nearest-node",
			testHashMap,
			"15",
			"6",
		},
		{
			"hash-nearest-node",
			testHashMap,
			"27",
			"8",
		},

		{
			"hash-overflow-cycle",
			testHashMap,
			"29",
			"2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.Get(tt.key); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
