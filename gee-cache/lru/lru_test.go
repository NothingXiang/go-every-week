package lru

import (
	"reflect"
	"testing"
)

func TestCache_Get(t *testing.T) {

	lru := NewCache(0, nil)

	lru.Add("key1", String("1234"))

	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields *Cache
		args   args
		wantV  Value
		wantOk bool
	}{
		{
			"key-exists",
			lru,
			args{key: "key1"},
			String("1234"),
			true,
		},
		{
			"key-not-exists",
			lru,
			args{key: "key2"},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotV, gotOk := tt.fields.Get(tt.args.key)
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("Get() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

type String string

func (d String) Len() int {
	return len(d)
}

func TestCache_RemoveOldest(t *testing.T) {
	keys := []string{"key1", "key2", "k3"}
	values := []string{"value1", "value2", "v3"}

	lru := NewCache(int64(len(keys[0]+keys[1]+values[0]+values[1])), nil)

	for i := 0; i < 3; i++ {
		lru.Add(keys[i], String(values[i]))
	}

	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}
