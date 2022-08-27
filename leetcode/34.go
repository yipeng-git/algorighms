package leetcode

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	index := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			index = mid
			break
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if index == -1 {
		return []int{-1, -1}
	}
	l, r = index, index
	for l > 0 {
		if l-1 >= 0 && nums[l-1] == target {
			l = l - 1
		} else {
			break
		}
	}
	for r < len(nums) {
		if r+1 < len(nums) && nums[r+1] == target {
			r += 1
		} else {
			break
		}
	}
	return []int{l, r}
}
