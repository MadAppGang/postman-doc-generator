package structtag

import (
	"testing"
)

func TestGetName(t *testing.T) {
	type args struct {
		tag string
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Get name of json tag",
			args{
				tag: `json:"id,omitempty"`,
				key: "json",
			},
			"id",
		},
		{
			"Get name of custom tag",
			args{
				tag: `abc:"name,omitempty"`,
				key: "abc",
			},
			"name",
		},
		{
			"Get empty string when key is missing in tag",
			args{
				tag: `json:"id,omitempty"`,
				key: "bson",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNameFromTag(tt.args.tag, tt.args.key); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
