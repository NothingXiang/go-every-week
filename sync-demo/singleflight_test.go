package sync_demo

import (
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestSingleMap_Get(t *testing.T) {
	type testStruct struct {
		name    string
		fields  *SingleMap
		args    string
		want    string
		wantErr bool
	}

	testSingleMap := NewSingleMap(mockSlowGet)
	tests := make([]testStruct, 0)

	// add 1000 goroutines to test get from singleMap with same key
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		tests = append(tests, testStruct{
			"test-sample-key",
			testSingleMap,
			"same-key",
			"same-key",
			false,
		})
	}

	for _, tt := range tests {
		go func(tt testStruct) {
			defer wg.Done()

			got, err := tt.fields.Get(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		}(tt)
	}

	wg.Wait()

	assert.Equal(t, cal, int64(1))
}

var cal int64

// slow getter mock ,log load number into cal
func mockSlowGet(key string) string {
	time.Sleep(10 * time.Second)
	log.Println("mockSlowGet,[key]:", key)
	atomic.AddInt64(&cal, 1)
	return key
}
