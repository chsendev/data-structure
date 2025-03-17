package leetcode_top100

import "testing"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := &ListNode{}

	p1 := list1
	p2 := list2

	t := newHead

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			t.Next = p1
			t = t.Next
			p1 = p1.Next
		} else {
			t.Next = p2
			t = t.Next
			p2 = p2.Next
		}
	}
	for p1 != nil {
		t.Next = p1
		t = t.Next
		p1 = p1.Next
	}
	for p2 != nil {
		t.Next = p2
		t = t.Next
		p2 = p2.Next
	}

	return newHead.Next
}

func TestMergeTwoLists(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 4}

	n4 := &ListNode{Val: 1}
	n5 := &ListNode{Val: 3}
	n6 := &ListNode{Val: 4}

	n1.Next = n2
	n2.Next = n3

	n4.Next = n5
	n5.Next = n6

	mergeTwoLists(n1, n4)
}
