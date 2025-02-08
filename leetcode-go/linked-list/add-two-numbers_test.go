package linked_list

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	fmt.Printf("%#+v", addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans *ListNode
	n1 := 0
	n2 := 0
	sum := 0
	carry := 0
	val := 0
	var t *ListNode
	for l1 != nil || l2 != nil {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
		}
		sum = n1 + n2 + carry

		val = sum % 10
		carry = sum / 10
		if ans == nil {
			ans = &ListNode{
				Val:  val,
				Next: nil,
			}
			t = ans
		} else {
			t.Next = &ListNode{
				Val:  val,
				Next: nil,
			}
			t = t.Next
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		t.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return ans
}
