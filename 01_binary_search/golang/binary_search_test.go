package main

import "testing"

func Test_binSearch(t *testing.T) {
	type args struct {
		list []int
		in   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				list: []int{1, 2, 3, 4, 5},
				in:   1,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				in:   10,
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				list: []int{1, 2, 3, 4, 5},
				in:   0,
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				in:   5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binSearch(tt.args.list, tt.args.in); got != tt.want {
				t.Errorf("binSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
