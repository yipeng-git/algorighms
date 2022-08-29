package leetcode

func quickSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	pivot := partition(nums, 0, len(nums)-1)
	quickSort(nums[:pivot])
	quickSort(nums[pivot+1:])
}

func partition(nums []int, low, high int) int {

	pivot := nums[high]
	i := low
	for j := low; j < high; j += 1 {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i += 1
		}
	}
	nums[i], nums[high] = nums[high], nums[i]

	return i
}
