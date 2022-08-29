package leetcode

import "math"

func jump(nums []int) int {
	jumps := make([]int, len(nums))
	for i := range jumps {
		if i >= 1 {
			jumps[i] = math.MaxInt16
		}
	}

	for i, _ := range jumps {
		next := i + nums[i]
		prev := i - nums[i]
		for j := i + 1; j <= next && j < len(nums); j += 1 {
			tmp := jumps[i] + 1
			if tmp < jumps[j] {
				jumps[j] = tmp
			}
		}
		for j := i - 1; j >= 0 && j >= prev; j -= 1 {
			tmp := jumps[i] + 1
			if tmp < jumps[j] {
				jumps[j] = tmp
			}
		}
	}
	return jumps[len(nums)-1]
}
