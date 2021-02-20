package main

import "testing"

func Test_subSequeue(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "fosh-fish",
			args: args {
				a: "fosh",
				b: "fish",
			},
			want: 3,
		},
		{
			name: "fosh-fort", 
			args: args {
				a: "fosh", 
				b: "fort",
			}, 
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subSequeue(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("subSequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}
