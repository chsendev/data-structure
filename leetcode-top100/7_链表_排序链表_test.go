package leetcode_top100

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {
	l1 := createListNode(5, 4, 3, 2, 1)
	l1 = sortList(l1)
	fmt.Println(l1)
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 快慢指针找出中位点
	// 1 2 3 4 5 6 7 8 9
	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 对右半部分进行归并排序
	r := mergeSort(slow.Next)
	// 链表判断结束的标志：末尾节点.next==null
	slow.Next = nil
	// 对左半部分进行归并排序
	l := mergeSort(head)
	return mergeList(l, r)
}

func mergeList(l, r *ListNode) *ListNode {
	// 临时头节点
	tmpHead := &ListNode{}
	p := tmpHead
	for l != nil && r != nil {
		if l.Val < r.Val {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}
	if l == nil {
		p.Next = r
	} else {
		p.Next = l
	}
	return tmpHead.Next
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestGetMaxLen(t *testing.T) {
	r1 := &TreeNode{Val: 1}
	r2 := &TreeNode{Val: 2}
	r3 := &TreeNode{Val: 3}
	r4 := &TreeNode{Val: 4}
	r5 := &TreeNode{Val: 5}
	r6 := &TreeNode{Val: 6}

	r1.Left = r2
	r1.Right = r3
	r2.Left = r4
	r2.Right = r5
	r4.Left = r6

	fmt.Println(getMaxLen(r1))
}

func getMaxLen(root *TreeNode) int {
	if root == nil {
		return 0
	}

	fmt.Println(root.Val)
	i1 := getMaxLen(root.Left)
	i2 := getMaxLen(root.Right)
	return max(i1, i2) + 1
}

func TestQuickSort(t *testing.T) {
	arr := []int{3, 5, 1, 7, 4, 2, 5, 9, 8}
	//partition(arr, 0, len(arr)-1, 2)

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	x := arr[l]
	m := partition(arr, l, r, x)
	quickSort(arr, l, m-1)
	quickSort(arr, m+1, r)
}

func partition(arr []int, l, r, x int) int {
	a := l // 指向大于x的第一位
	xi := 0
	for i := l; i <= r; i++ {
		if arr[i] <= x {
			arr[i], arr[a] = arr[a], arr[i]
			if arr[a] == x {
				xi = a
			}
			a++
		}
	}

	arr[xi], arr[a-1] = arr[a-1], arr[xi]
	return a - 1
}

func TestName(t *testing.T) {
	//fmt.Println(foo())
	arr := []int{5, 4, 3, 2, 1}
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[0])
		arr[0], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[0]
		heapify(arr, 0, len(arr)-i-1)
	}
}

func heapInsert(arr []int, i int) {
	parent := (i - 1) / 2
	for parent >= 0 {
		if arr[i] > arr[parent] {
			arr[i], arr[parent] = arr[parent], arr[i]
			i = parent
			parent = (i - 1) / 2
		} else {
			break
		}
	}
}

func heapify(arr []int, i, size int) {
	left := i*2 + 1
	for left < size {
		right := left + 1
		most := left
		if right < size && arr[right] > arr[left] {
			most = right
		}
		if arr[most] > arr[i] {
			arr[most], arr[i] = arr[i], arr[most]
			i = most
			left = i*2 + 1
		} else {
			break
		}
	}
}

//func foo() error {
//	e := errors.New("D")
//	defer func() {
//		fmt.Println(e)
//		e = errors.New("A")
//	}()
//	defer func() {
//		fmt.Println(e)
//		e = errors.New("B")
//	}()
//	e = errors.New("C")
//	return e
//}
