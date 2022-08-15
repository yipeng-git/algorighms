package leetcode

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

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, _ := range nums {
		if other, ok := m[target-nums[i]]; ok {
			return []int{other, i}
		}
	}
	return nums
}
