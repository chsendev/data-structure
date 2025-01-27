package data_structure

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestName(t *testing.T) {
	lists := make([][]*ListNode, 0)
	a1 := &ListNode{Val: 1, Next: nil}
	a2 := &ListNode{Val: 4, Next: nil}
	a3 := &ListNode{Val: 5, Next: nil}
	a1.Next = a2
	a2.Next = a3

	b1 := &ListNode{Val: 1, Next: nil}
	b2 := &ListNode{Val: 3, Next: nil}
	b3 := &ListNode{Val: 4, Next: nil}
	b1.Next = b2
	b2.Next = b3

	lists = append(lists, []*ListNode{
		a1, a2, a3,
	})
	lists = append(lists, []*ListNode{
		b1, b2, b3,
	})
	//mergeKLists(lists)
	//newLists := make([]*ListNode, 0)

	//for _, list := range lists {
	//	for _, node := range list {
	//		newLists = append(newLists, node)
	//	}
	//}
	//
	//newLists := make([]*ListNode, 0)
}

//func mergeKLists(lists []*Lis[][]tNode) *ListNode {
//
//}
