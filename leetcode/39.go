package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	combinationSumBackTracking(candidates, target, 0, 0, []int{}, &res)
	return res
}

func combinationSumBackTracking(candidates []int, target int, sum int, idx int, arr []int, res *[][]int) {
	if sum == target {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*res = append(*res, tmp)
		return
	}
	if idx >= len(candidates) || sum > target {
		return
	}
	for i := idx; i < len(candidates); i += 1 {
		arr = append(arr, candidates[i])
		combinationSumBackTracking(candidates, target, sum+candidates[i], i, arr, res)
		arr = arr[:len(arr)-1]
	}
}
