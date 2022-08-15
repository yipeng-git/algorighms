package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{"aabc"},
			want: 3,
		},
		{
			name: "",
			args: args{"abcabcbb"},
			want: 3,
		},
		{
			name: "",
			args: args{"bbbbbb"},
			want: 1,
		},
		{
			name: "",
			args: args{"pwwkew"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
