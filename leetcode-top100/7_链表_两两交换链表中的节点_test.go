package leetcode_top100

import "testing"

func swapPairs(head *ListNode) *ListNode {
	head = &ListNode{Next: head}

	p := head
	for p.Next != nil {
		if p.Next.Next == nil {
			break
		}

		t1 := p.Next
		t2 := p.Next.Next
		t1.Next = t2.Next
		t2.Next = t1
		p.Next = t2

		p = p.Next.Next
	}
	return head.Next
}

func TestSwapPairs(t *testing.T) {
	l1 := createListNode(1, 2, 3, 4)
	swapPairs(l1)
}
