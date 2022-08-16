package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		curr *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				curr: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			},
			want: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.curr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseAndPop(t *testing.T) {
	type args struct {
		curr *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				curr: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
				n:    1,
			},
			want: &ListNode{Val: 3, Next: &ListNode{Val: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseAndPop(tt.args.curr, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseAndPop() = %v, want %v", got, tt.want)
			}
		})
	}
}
