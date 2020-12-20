package gee

import (
	"reflect"
	"testing"
)

func Test_parsePattern(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name      string
		args      args
		wantParts []string
	}{
		{
			"empty",
			args{pattern: ""},
			[]string{},
		},
		{
			"ignore-after-*",
			args{pattern: "/api/*filePath/name"},
			[]string{"api", "*filePath"},
		},
		{
			"only-one-/",
			args{pattern: "/"},
			[]string{},
		},
		{
			"common",
			args{pattern: "/api/find/name"},
			[]string{"api", "find", "name"},
		},
		{
			"*",
			args{pattern: "/api/find/*filePath"},
			[]string{"api", "find", "*filePath"},
		},
		{
			"match-param",
			args{pattern: "/api/find/:name"},
			[]string{"api", "find", ":name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotParts := parsePattern(tt.args.pattern); !reflect.DeepEqual(gotParts, tt.wantParts) {
				t.Errorf("parsePattern() = %v, want %v", gotParts, tt.wantParts)
			}
		})
	}
}
