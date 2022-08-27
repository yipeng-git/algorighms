package leetcode

import "sort"

func nextPermutation(nums []int) {
	sortIndex := -1
out:
	for i := len(nums) - 2; i >= 0; i -= 1 {
		for j := len(nums) - 1; j > i; j -= 1 {
			if nums[i] < nums[j] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
				sortIndex = i + 1
				break out
			}
		}
	}
	if sortIndex == -1 {
		sortIndex = 0
	}
	sort.Ints(nums[sortIndex:])
}
