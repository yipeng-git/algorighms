package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	index := 0
	for index < len(preorder) {
		if inorder[index] == preorder[0] {
			break
		}
		index += 1
	}
	node := &TreeNode{Val: preorder[0]}
	node.Left = buildTree(preorder[1:index+1], inorder[:index])
	node.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return node
}
