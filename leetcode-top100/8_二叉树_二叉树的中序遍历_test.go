package leetcode_top100

import (
	"fmt"
	"testing"
)

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	r1 := inorderTraversal(root.Left)
	r1 = append(r1, root.Val)
	r2 := inorderTraversal(root.Right)
	return append(r1, r2...)
}

func TestInorderTraversal(t *testing.T) {
	r1 := &TreeNode{Val: 1}
	r2 := &TreeNode{Val: 2}
	r3 := &TreeNode{Val: 3}
	//r4 := &TreeNode{Val: 4}
	//r5 := &TreeNode{Val: 5}

	r1.Right = r2
	r2.Left = r3
	//r2.Left = r4
	//r2.Right = r5

	fmt.Println(inorderTraversal(r1))

}
