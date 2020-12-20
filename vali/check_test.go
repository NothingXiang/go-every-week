package vali

import (
	"testing"
)

func TestCheck(t *testing.T) {

	tests := []struct {
		name    string
		args    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", nil, true},
		{"", TstStruct1{Const: "others"}, true},
		{"", TstStruct1{Const: "const"}, false},
		{"", User{}, false},
		{"", User{FirstName: "1"}, false},
		{"", User{FirstName: "1", Email: "nothing@m.scnu.edu.cn"}, false},
		{"", User{FirstName: "1", Email: "nothing@m.scnu.edu.cn", FavourColor: "#000-"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Check(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
