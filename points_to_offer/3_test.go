package points_to_offer

import "testing"

func TestSolution3(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "no",
			args: args{[]int{0, 1, 2}},
			want: -1,
		},
		{
			name: "normal1",
			args: args{[]int{2, 1, 2}},
			want: 2,
		},
		{
			name: "normal2",
			args: args{[]int{2, 3, 1, 0, 2, 5}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution3(tt.args.slice); got != tt.want {
				t.Errorf("Solution3() = %v, want %v", got, tt.want)
			}
		})
	}
}
