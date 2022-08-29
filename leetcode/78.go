package leetcode

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	subsetsHelper(nums, []int{}, 0, &res)
	return res
}

func subsetsHelper(nums []int, curr []int, index int, res *[][]int) {
	if index >= len(nums) {
		*res = append(*res, curr)
		return
	}
	subsetsHelper(nums, curr, index+1, res)

	tmp := make([]int, len(curr))
	copy(tmp, curr)
	tmp = append(tmp, nums[index])
	subsetsHelper(nums, tmp, index+1, res)
}
