package leetcode

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				nums: []int{3, 2, 1},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 8, 1, 4, 5, 6, 7, 2},
			},
		},
		{
			name: "1",
			args: args{
				nums: []int{5, 3, 6, 78, 3, 1, 7, 8, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
