package leetcode

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	can := make([]int, len(nums))
	for i, _ := range nums {
		if i > 0 && can[i] == 0 {
			return false
		}
		next := nums[i] + i
		tmp := can[i] + 1
		for j := i + 1; j <= next && j < len(nums); j += 1 {
			if can[j] == 0 {
				can[j] = tmp
			} else {
				if can[j] > tmp {
					can[j] = tmp
				}
			}
		}
	}
	return can[len(nums)-1] > 0
}
