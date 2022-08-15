package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	//m := map[int]bool{}
	//newNums := []int{}
	////for _, n := range nums {
	////	if _, ok := m[n]; ok {
	////		continue
	////	}
	////	newNums = append(newNums, n)
	////	m[n] = true
	////}
	//nums = newNums
	res := [][]int{}
	for i := 0; i < len(nums)-2; i += 1 {
		partial := tSum(nums[i+1:], -nums[i])
		if len(partial) > 0 {
			for _, p := range partial {
				res = append(res, append([]int{nums[i]}, p...))

			}

		}
	}
	return res
}

func tSum(nums []int, target int) [][]int {
	var res [][]int
	m := map[int]int{}
	for i := 0; i < len(nums); i += 1 {
		if pos, ok := m[target-nums[i]]; ok {
			existed := false
			for _, already := range res {
				if (already[0] == nums[pos] && already[1] == nums[i]) || already[1] == nums[pos] && already[0] == nums[i] {
					existed = true
				}
			}
			if existed {
				continue
			}
			res = append(res, []int{nums[pos], nums[i]})
		}
		m[nums[i]] = i
	}
	return res
}
