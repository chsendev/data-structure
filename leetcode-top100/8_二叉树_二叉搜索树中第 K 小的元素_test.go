package leetcode_top100

func kthSmallest(root *TreeNode, k int) int {
	var list []int
	kthSmallestList(root, &list)
	return list[k-1]
}

func kthSmallestList(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	kthSmallestList(root.Left, arr)
	*arr = append(*arr, root.Val)
	kthSmallestList(root.Right, arr)
}
