package main

import (
	"reflect"
	"testing"
)

func Test_factor(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{n: 6},
			want: []int{2, 3},
		},
		{
			name: "test2",
			args: args{n: 8},
			want: []int{2, 2, 2},
		},
		{
			name: "test3",
			args: args{n: 90},
			want: []int{2, 3, 3, 5},
		},
		{
			name: "test4",
			args: args{n: 7},
			want: []int{7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factor(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factor() = %v, want %v", got, tt.want)
			}
		})
	}
}
