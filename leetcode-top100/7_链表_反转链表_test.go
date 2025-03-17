package leetcode_top100

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	p := head
	var newHead *ListNode
	for p != nil {
		t := p
		p = p.Next
		if newHead == nil {
			newHead = t
			newHead.Next = nil
		} else {
			t.Next = newHead
			newHead = t
		}
	}
	return newHead
}

func TestReverseList(t *testing.T) {

}
