package leetcode

import (
	"fmt"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			fmt.Println(tt.args.matrix)
		})
	}
}
