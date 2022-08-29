package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	currLevel := []*TreeNode{root}
	var currVals []int
	var nextLevel []*TreeNode
	var res [][]int
	for len(currLevel) > 0 {
		currNode := currLevel[0]
		currVals = append(currVals, currNode.Val)
		if currNode.Left != nil {
			nextLevel = append(nextLevel, currNode.Left)
		}
		if currNode.Right != nil {
			nextLevel = append(nextLevel, currNode.Right)
		}
		currLevel = currLevel[1:]
		if len(currLevel) == 0 {
			currLevel = nextLevel
			nextLevel = nil
			res = append(res, currVals)
			currVals = nil
		}
	}
	return res
}
