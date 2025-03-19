package leetcode_top100

func sortedArrayToBST(nums []int) *TreeNode {
	return toSortedArrayToBST(nums, 0, len(nums)-1)
}

func toSortedArrayToBST(nums []int, l, r int) *TreeNode {
	if r < l {
		return nil
	}
	mid := l + (r-l)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = toSortedArrayToBST(nums, l, mid-1)
	root.Right = toSortedArrayToBST(nums, mid+1, r)
	return root
}
