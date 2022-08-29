package leetcode

func sortColors(nums []int) {
	//cnt := []int{0, 0, 0}
	end := len(nums) - 1
	for i := 0; i <= end; i += 1 {
		if nums[i] == 0 {
			continue
		}
		curr := nums[i]
		j := end
		found := false
		for ; j > i; j -= 1 {
			if nums[j] != 2 && nums[j] < curr {
				found = true
				break
			}
		}
		if found {
			nums[i] = nums[j]
			nums[j] = curr
			i -= 1
		}
	}
}
