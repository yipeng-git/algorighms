package leetcode

func maxArea(height []int) int {
	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	minInt := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	l := 0
	r := len(height) - 1
	var res int

	for l < r {
		res = maxInt(minInt(height[l], height[r])*(r-l), res)
		if height[l] < height[r] {
			l += 1
		} else {
			r -= 1
		}
	}
	return res
}
