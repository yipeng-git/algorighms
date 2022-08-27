package leetcode

func permute(nums []int) [][]int {
	res := [][]int{}
	permuteBacktracking([]int{}, nums, &res)
	return res
}

func permuteBacktracking(path, nums []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, path)
		return
	}
	for i, num := range nums {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		permuteBacktracking(append(pathCopy, num), append(numsCopy[:i], numsCopy[i+1:]...), result)
	}
}
