package gee_cache

import (
	"fmt"
	"log"
	"testing"
)

var DB = map[string]string{
	"Nothing": "Xiang",
	"Sin":     "Xiang",
	"Vincent": "Xing",
}

func TestGroup_Get(t *testing.T) {

	loadCounts := make(map[string]int, len(DB))

	mockDBGetter := func(key string) ([]byte, error) {
		log.Println("[SlwDB] search key", key)
		if value, ok := DB[key]; ok {
			loadCounts[key] += 1
			return []byte(value), nil
		}
		return nil, fmt.Errorf("%s not exists", key)
	}

	groupName := "test-scores"
	NewGroup(groupName, 2<<10, GetterFunc(mockDBGetter))

	for k, v := range DB {
		if view, err := GetGroup(groupName).Get(k); err != nil || view.String() != v {
			t.Fatalf("failed to get value of %s", k)
		} // load from callback function
		if _, err := GetGroup(groupName).Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("safelyCache %s miss", k)
		} // safelyCache hit
	}

	if view, err := GetGroup(groupName).Get("unknown"); err == nil {
		t.Fatalf("the value of unknow should be empty, but %s got", view)
	}
}
