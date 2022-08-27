package leetcode

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	index := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			index = mid
			break
		}
		if nums[mid] >= nums[l] {
			if target >= nums[l] && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return index
}
