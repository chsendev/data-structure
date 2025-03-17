package leetcode_top100

import (
	"fmt"
	"testing"
)

func isPalindrome(head *ListNode) bool {
	p := head
	var reverseHead *ListNode
	for p != nil {
		t := &ListNode{Val: p.Val}
		p = p.Next
		if reverseHead == nil {
			reverseHead = t
			reverseHead.Next = nil
		} else {
			t.Next = reverseHead
			reverseHead = t
		}
	}

	p = head
	p2 := reverseHead
	for p != nil && p2 != nil {
		if p.Val != p2.Val {
			return false
		}
		p = p.Next
		p2 = p2.Next
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 1}
	n3 := &ListNode{Val: 2}
	n4 := &ListNode{Val: 1}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	fmt.Println(isPalindrome(n1))
}
