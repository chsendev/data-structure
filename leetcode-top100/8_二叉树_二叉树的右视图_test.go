package leetcode_top100

import (
	"container/list"
	"fmt"
	"testing"
)

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{root.Val}
	queue := list.New()
	if root.Left != nil {
		queue.PushBack(root.Left)
	}
	if root.Right != nil {
		queue.PushBack(root.Right)
	}

	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			ele := queue.Front()
			t := ele.Value.(*TreeNode)
			if t != nil {
				if t.Left != nil {
					queue.PushBack(t.Left)
				}
				if t.Right != nil {
					queue.PushBack(t.Right)
				}
				if i == size-1 {
					res = append(res, t.Val)
				}
			}
			queue.Remove(ele)
		}
	}
	return res
}

func createTreeNode(val ...int) *TreeNode {
	return createTree(val, 0)
}

func createTree(val []int, i int) *TreeNode {
	if i >= len(val) {
		return nil
	} else if val[i] == -1 {
		return nil
	}
	root := &TreeNode{Val: val[i]}
	root.Left = createTree(val, 2*i+1)
	root.Right = createTree(val, 2*i+2)

	return root
}

func TestRightSideView(t *testing.T) {
	//root := createTreeNode(1, 2, 3, -1, 5, -1, 4)
	root := createTreeNode(1, 2)
	fmt.Println(root)
	fmt.Println(rightSideView(root))
}
