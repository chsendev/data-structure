package leetcode_top100

import (
	"container/list"
	"fmt"
	"testing"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	l := list.New()
	res := [][]int{
		{root.Val},
	}

	l.PushBack(root.Left)
	l.PushBack(root.Right)

	for l.Len() > 0 {
		var res2 []int
		size := l.Len()
		for i := 0; i < size; i++ {
			front := l.Front()
			if front != nil {
				n := front.Value.(*TreeNode)
				if n != nil {
					res2 = append(res2, n.Val)
					l.PushBack(n.Left)
					l.PushBack(n.Right)
				}
				l.Remove(front)
			}
		}
		if len(res2) > 0 {
			res = append(res, res2)
		}
	}
	return res
}

func TestLevelOrder(t *testing.T) {
	r1 := &TreeNode{Val: 1}
	r2 := &TreeNode{Val: 2}
	r3 := &TreeNode{Val: 3}
	r4 := &TreeNode{Val: 4}
	r5 := &TreeNode{Val: 5}

	r1.Left = r2
	r1.Right = r3
	r2.Left = r4
	r2.Right = r5

	fmt.Println(levelOrder(r1))

}
