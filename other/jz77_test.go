package other

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrint(t *testing.T) {
	type args struct {
		pRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Print(tt.args.pRoot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPath(t *testing.T) {
	l := &TreeNode{Val: 2}
	l.Left = &TreeNode{Val: 4}
	l.Right = &TreeNode{Val: 5}
	l.Right.Right = &TreeNode{Val: -1}
	r := &TreeNode{Val: 3}
	r.Left = &TreeNode{Val: 4}
	r.Right = &TreeNode{Val: 3}
	in := &TreeNode{Val: 1}
	in.Left = l
	in.Right = r
	res := FindPath(in, 6)
	if res != 3 {
		t.Errorf("wrong %v, expected %v", res, 3)
	}
}

func TestMinRotate(t *testing.T) {
	res := minNumberInRotateArray([]int{3, 4, 5, 1, 2})
	if res != 1 {
		t.Errorf("wrong %v, expected 1", res)
	}
	res = minNumberInRotateArray([]int{2, 2, 2, 1, 2})
	if res != 1 {
		t.Errorf("wrong %v, expected 1", res)
	}
}

func TestPermutation(t *testing.T) {
	res := Permutation("ab")
	fmt.Println(res)
}

func TestFindNthDigit(t *testing.T) {
	a := findNthDigit(500000000)
	fmt.Println(a)
}

func TestFindGreatestSumOfSubArray2(t *testing.T) {
	res := FindGreatestSumOfSubArray2([]int{1, -2, 3, 10, -4, 7, 2, -5})
	fmt.Println(res)
}

func Test_jumpFloorII(t *testing.T) {
	res := jumpFloorII(3)
	fmt.Println(res)
}

func Test_solve(t *testing.T) {
	//res := solve("12")
	//fmt.Println(res)
	res = solve("31717126241541717")
	fmt.Println(res)
}

func Test_insert(t *testing.T) {
	Insert(1)
	Insert(2)
	Insert(1)
	fmt.Println(arr)
	Insert(2)
	fmt.Println(arr)
	mid := GetMedian()
	fmt.Println(mid)
}

func TestAdd(t *testing.T) {
	res := Add(1, 1)
	if res != 2 {
		t.Errorf("wrong %v, expected %v", res, 2)
	}
}

func TestNumberOf1(t *testing.T) {
	NumberOf1(-1)
}
