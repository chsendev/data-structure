package leetcode_top100

import (
	"fmt"
	"testing"
)

var pre *TreeNode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := isValidBST(root.Left)
	if pre != nil {
		if root.Val <= pre.Val {
			return false
		}
	}
	pre = root
	right := isValidBST(root.Right)
	return left && right
}

func TestIsValidBST(t *testing.T) {
	r1 := &TreeNode{Val: 0}
	//r5 := &TreeNode{Val: 5}
	//r4 := &TreeNode{Val: 6}
	//r3 := &TreeNode{Val: 3}
	//r6 := &TreeNode{Val: 6}
	//
	//r5.Left = r1
	//r5.Right = r4
	//r4.Left = r3
	//r4.Right = r6
	fmt.Println(isValidBST(r1))
}
