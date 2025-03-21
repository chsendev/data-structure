package leetcode_top100

import (
	"fmt"
	"testing"
)

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func TestHasCycle(t *testing.T) {
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: 4}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2

	fmt.Println(hasCycle(n1))
}
