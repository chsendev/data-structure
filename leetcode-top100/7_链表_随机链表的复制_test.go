package leetcode_top100

import "testing"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)

	p := head
	newHead := &Node{}
	t := newHead
	for p != nil {
		m[p] = &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		p = p.Next
	}

	p = head
	for p != nil {
		t.Next = m[p]
		if t.Next != nil {
			t.Next.Next = m[p.Next]
			t.Next.Random = m[p.Random]
			t = t.Next
		}
		p = p.Next
	}

	return newHead.Next
}

func TestCopyRandomList(t *testing.T) {

}
