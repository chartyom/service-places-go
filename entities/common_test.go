package entities

import "testing"

func TestCommon_StringError(t *testing.T) {
	type args struct {
		v map[string]string
	}
	tests := []struct {
		name string
		c    *Common
		args args
		want string
	}{
		{
			name: "First",
			c:    &Common{},
			args: args{
				v: map[string]string{
					"first":  "1",
					"second": "2",
					"third":  "3",
				},
			},
			want: "first:1;second:2;third:3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.StringError(tt.args.v); got != tt.want {
				t.Errorf("Common.StringError() = %v, want %v", got, tt.want)
			}
		})
	}
}
