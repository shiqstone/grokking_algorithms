package main

import (
	"reflect"
	"testing"
)

func TestGetCoverStations(t *testing.T) {
	type args struct {
		stations   map[string][]string
		statesNeed []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args {
				stations: map[string][]string {
					"kone":[]string{"id", "nv", "ut"},
					"ktwo":[]string{"wa", "id", "mt"},
					"kthree":[]string{"or", "nv", "ca"},
					"kfour":[]string{"nv", "ut"},
					"kfive":[]string{"ca", "az"},
				},
				statesNeed: []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			},
			want: []string{"kone", "ktwo", "kthree", "kfive"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCoverStations(tt.args.stations, tt.args.statesNeed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoverStations() = %v, want %v", got, tt.want)
			}
		})
	}
}
