package main

import "testing"

func Test_substr(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "hish-vista",
			args: args {
				a: "vista",
				b: "hish",
			},
			want: "is",
		},
		{
			name: "hish-fish", 
			args: args {
				a: "fish", 
				b: "hish",
			}, 
			want: "ish",
		},
		{
			name: "fish-fosh", 
			args: args {
				a: "fish", 
				b: "hosh",
			}, 
			want: "sh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substr(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("substr() = %v, want %v", got, tt.want)
			}
		})
	}
}
