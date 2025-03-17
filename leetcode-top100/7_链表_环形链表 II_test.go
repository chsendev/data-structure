package leetcode_top100

import (
	"fmt"
	"testing"
)

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	m := make(map[*ListNode]struct{})
	p := head
	for p != nil {
		_, ok := m[p]
		if ok {
			return p
		}
		m[p] = struct{}{}

		p = p.Next
	}
	return nil
}

func TestDetectCycle(t *testing.T) {
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: 4}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2

	fmt.Println(detectCycle(n1))
}
