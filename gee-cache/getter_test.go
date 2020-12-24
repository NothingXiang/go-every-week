package gee_cache

import (
	"reflect"
	"testing"
)

func TestGetterFunc_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		f       GetterFunc
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"simple-getter",
			func(key string) ([]byte, error) {
				return []byte(key), nil
			},
			args{key: "key"},
			[]byte("key"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
