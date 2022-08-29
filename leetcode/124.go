package leetcode

import "math"

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var helper func(node *TreeNode) int
	max := math.MinInt
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			if node.Val > max {
				max = node.Val
			}
			if node.Val > 0 {
				return node.Val
			}
			return 0
		}
		left := helper(node.Left)
		right := helper(node.Right)
		sum := left + right + node.Val
		if sum > max {
			max = sum
		}
		if left > right {
			if left+node.Val > 0 {
				return left + node.Val
			} else {
				return 0
			}
		} else {
			if right+node.Val > 0 {
				return right + node.Val
			} else {
				return 0
			}
		}
	}
	helper(root)
	return max
}
