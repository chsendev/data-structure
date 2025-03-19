package leetcode_top100

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return cmp(root.Left, root.Right)
}

func cmp(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return cmp(left.Left, right.Right) && cmp(left.Right, right.Left)
}
