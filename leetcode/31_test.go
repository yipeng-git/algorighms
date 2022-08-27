package leetcode

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	input := []int{1, 2, 3}
	nextPermutation(input)
	fmt.Println(input)
	input = []int{1, 1, 3}
	nextPermutation(input)
	fmt.Println(input)
}
