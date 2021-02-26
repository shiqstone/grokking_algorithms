package main

import "testing"

func Test_generateDiff(t *testing.T) {
	type args struct {
		src []string
		dst []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test_base_chart",
			args: args{
				src: []string{"A", "B", "C", "A", "B", "B", "A"},
				dst: []string{"C", "B", "A", "B", "A", "C"},
			},
		},
		{
			name: "test_chart_2",
			args: args{
				src: []string{"A", "B", "C"},
				dst: []string{"D", "E", "F"},
			},
		},
		{
			name: "test_char3",
			args: args{
				src: []string{"A", "B", "C", "A", "B", "B", "A"},
				dst: []string{"A", "B"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateDiff(tt.args.src, tt.args.dst)
		})
	}
}
