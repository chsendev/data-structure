package leetcode_top100

func diameterOfBinaryTree(root *TreeNode) int {
	return maxDepth2(root) - 1
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth2(root.Left), maxDepth2(root.Right)) + 1
}
