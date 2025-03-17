package leetcode_top100

import "testing"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2

	addNum := 0
	newHead := &ListNode{}
	p3 := newHead
	for p1 != nil && p2 != nil {
		t := p1.Val + p2.Val + addNum
		addNum = t / 10
		t %= 10
		p3.Next = &ListNode{Val: t}
		p3 = p3.Next

		p1 = p1.Next
		p2 = p2.Next
	}

	for p1 != nil {
		t := p1.Val + addNum
		addNum = t / 10
		t %= 10
		p3.Next = &ListNode{Val: t}
		p3 = p3.Next

		p1 = p1.Next
	}
	for p2 != nil {
		t := p2.Val + addNum
		addNum = t / 10
		t %= 10
		p3.Next = &ListNode{Val: t}
		p3 = p3.Next

		p2 = p2.Next
	}

	if addNum > 0 {
		p3.Next = &ListNode{Val: addNum}
	}

	return newHead.Next
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := createListNode(9, 9, 9, 9, 9, 9, 9)
	l2 := createListNode(9, 9, 9, 9)
	addTwoNumbers(l1, l2)
}

func createListNode(vals ...int) *ListNode {
	h := &ListNode{}
	t := h
	for _, val := range vals {
		t.Next = &ListNode{Val: val}
		t = t.Next
	}

	return h.Next
}
