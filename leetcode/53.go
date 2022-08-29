package leetcode

import "math"

func maxSubArray(nums []int) int {
	//start := 0
	//end := 0
	curMax := nums[0]
	sum := make([]int, len(nums))
	for i, _ := range sum {
		sum[i] = math.MinInt
	}
	sum[0] = nums[0]
	for i := 1; i < len(nums); i += 1 {
		tmp := nums[i] + sum[i-1]
		if tmp > nums[i] {
			sum[i] = tmp
		} else {
			sum[i] = nums[i]
		}
		if sum[i] > curMax {
			curMax = sum[i]
		}
	}
	return curMax

}
