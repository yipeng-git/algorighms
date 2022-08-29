package leetcode

import "sort"

//func twoSum(nums []int, target int) []int {
//	a := make(map[int]int, len(nums))
//	for i, _ := range nums {
//		if other, ok := a[target-nums[i]]; ok {
//			return []int{other, i}
//		}
//		a[nums[i]] = i
//	}
//	return []int{}
//}

//func twoSum(nums []int, target int) []int {
//	m := map[int]int{}
//	for i, _ := range nums {
//		if other, ok := m[target-nums[i]]; ok {
//			return []int{other, i}
//		}
//	}
//	return nums
//}

func twoSum(nums []int, target int) []int {
	sort.Ints(nums)
	for i, j := 0, 1; j < len(nums); {
		if i == j {
			j += 1
		}
		curr := nums[i] + nums[j]
		if curr == target {
			return []int{i, j}
		} else if curr < target {
			j += 1
		} else {
			i += 1
		}
	}
	return []int{}
}
