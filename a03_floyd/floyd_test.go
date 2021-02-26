package main

import (
	"reflect"
	"testing"
)

func Test_floyd(t *testing.T) {
	type args struct {
		graph map[string]map[string]int
		start string
		end   string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				graph: map[string]map[string]int{
					"start": map[string]int{
						"a": 6,
						"b": 2,
					},
					"a": map[string]int{
						"fin": 1,
					},
					"b": map[string]int{
						"a":   3,
						"fin": 5,
					},
					"fin": map[string]int{},
				},
				start: "start",
				end:   "fin",
			},
			wantRes: []string{
				"start",
				"b",
				"a",
				"fin",
			},
		},
		{
			name: "test2",
			args: args{
				graph: map[string]map[string]int{
					"1": map[string]int{
						"2": 2,
						"3": 3,
						"4": 6,
					},
					"2": map[string]int{
						"5": 4,
						"6": 6,
					},
					"3": map[string]int{
						"4": 2,
					},
					"4": map[string]int{
						"5": 1,
						"6": 3,
					},
					"5": map[string]int{},
					"6": map[string]int{},
				},
				start: "1",
				end:   "6",
			},
			wantRes: []string{
				"1",
				"2",
				"6",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := floyd(tt.args.graph, tt.args.start, tt.args.end); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("floyd() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
