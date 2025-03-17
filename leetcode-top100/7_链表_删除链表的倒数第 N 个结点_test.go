package leetcode_top100

import "testing"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = &ListNode{Next: head}
	fast := head
	slow := head

	i := 0
	for fast.Next != nil {
		if i >= n {
			slow = slow.Next
		}

		i++
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return head.Next
}

func TestRemoveNthFromEnd(t *testing.T) {
	l1 := createListNode(1, 2, 3, 4, 5)
	removeNthFromEnd(l1, 2)
}
