package leetcode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	left := root.Left
	root.Left = nil
	flatten(root.Right)
	right := root.Right
	root.Right = nil
	flattenAppend(root, left)
	flattenAppend(root, right)
}

func flattenAppend(base, tail *TreeNode) {
	for base.Right != nil {
		base = base.Right
	}
	base.Right = tail
}
