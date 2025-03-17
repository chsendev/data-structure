package leetcode_top100

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	p1 := headA
	for p1 != nil {
		m[p1] = struct{}{}
		p1 = p1.Next
	}

	p2 := headB
	for p2 != nil {
		_, ok := m[p2]
		if ok {
			return p2
		}
		p2 = p2.Next
	}
	return nil
}
