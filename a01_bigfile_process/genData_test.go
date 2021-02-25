package main

import "testing"

func Test_genMd5Data(t *testing.T) {
	type args struct {
		limit int
		path  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test data",
			args: args{
				limit: 1000000,
				path:  "d:/tmp/bigMd5TestData.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			genMd5Data(tt.args.limit, tt.args.path)
		})
	}
}

func Test_genRandNumData(t *testing.T) {
	type args struct {
		limit int
		path  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test data",
			args: args{
				limit: 1000000,
				path:  "d:/tmp/bigRandNumTestData.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			genRandNumData(tt.args.limit, tt.args.path)
		})
	}
}
