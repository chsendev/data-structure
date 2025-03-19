package leetcode_top100

import (
	"fmt"
	"testing"
)

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	l, r := root.Left, root.Right
	flatten(root.Left)
	flatten(root.Right)
	if l != nil {
		t := l
		for t.Right != nil { // 找到左子树的最后一个节点
			t = t.Right
		}
		t.Right = r
		root.Right = l
	} else {
		root.Right = r
	}
	root.Left = nil
}

func TestFlatten(t *testing.T) {
	root := createTreeNode(1, 2, 5, 3, 4, -1, 6)
	flatten(root)
	fmt.Println(root)
}
