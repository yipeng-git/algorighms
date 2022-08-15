package leetcode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				l1: &ListNode{Val: 8},
				l2: &ListNode{Val: 9},
			},
			want: &ListNode{Val: 7, Next: &ListNode{Val: 1}},
		},
		{
			name: "",
			args: args{
				l1: &ListNode{Val: 9},
				l2: &ListNode{Val: 9, Next: &ListNode{Val: 9}},
			},
			want: &ListNode{Val: 8, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
