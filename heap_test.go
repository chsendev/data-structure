package data_structure

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

type MaxHeap struct {
	Arr []int
}

func Create(arr []int) *MaxHeap {
	var h MaxHeap
	// 找到最后一颗子树
	for parent := (len(arr) - 1) / 2; parent >= 0; parent-- {
		var child int
		for i := parent; i*2+1 < len(arr); i = child {
			child = i*2 + 1
			if child+1 < len(arr) && arr[child] < arr[child+1] {
				child = child + 1
			}
			if arr[child] > arr[parent] {
				arr[parent], arr[child] = arr[child], arr[parent]
			}
		}
	}
	h.Arr = arr
	return &h
}

func (m *MaxHeap) Insert(item int) {
	m.Arr = append(m.Arr, item)
	var parent int
	for child := len(m.Arr) - 1; (child-1)/2 >= 0; child = parent {
		// 找到父节点
		parent = (child - 1) / 2
		// 是否比父节点大，是的话则交换
		if m.Arr[child] > m.Arr[parent] {
			m.Arr[child], m.Arr[parent] = m.Arr[parent], m.Arr[child]
		} else {
			break
		}
	}
}

func (m *MaxHeap) Insert2(item int) {
	m.Arr = append(m.Arr, item)
	m.Recursion(len(m.Arr) - 1)

}
func (m *MaxHeap) Recursion(child int) {
	if (child-1)/2 < 0 {
		return
	}
	parent := (child - 1) / 2
	if m.Arr[child] > m.Arr[parent] {
		m.Arr[child], m.Arr[parent] = m.Arr[parent], m.Arr[child]
		m.Recursion(parent)
	} else {
		return
	}
}

func (m *MaxHeap) Delete() int {
	maxEle := m.Arr[0]
	m.Arr[0] = m.Arr[len(m.Arr)-1]
	m.Arr = m.Arr[:len(m.Arr)-1]
	var child int
	for parent := 0; parent*2+1 < len(m.Arr); parent = child {
		child = parent*2 + 1
		if parent*2+2 < len(m.Arr) && m.Arr[child] < m.Arr[parent*2+2] {
			child = parent*2 + 2
		}
		if m.Arr[child] > m.Arr[parent] {
			m.Arr[child], m.Arr[parent] = m.Arr[parent], m.Arr[child]
		} else {
			break
		}
	}
	return maxEle
}

func Print(arr []int) {
	d := Deep(arr)
	for i := 0; i < d; i++ {
		layer := int(math.Pow(2, float64(i)))
		first := int(math.Pow(2, float64(i))) - 1
		var swap string
		for j := i; j < d; j++ {
			swap += "\t"
		}
		fmt.Print(swap)
		for j := first; j < layer+first && j < len(arr); j++ {
			fmt.Printf(strconv.Itoa(int(arr[j])))
			swap = ""
			for k := i; k < d; k++ {
				swap += "  "
			}
			fmt.Print(swap)
		}
		fmt.Println()
	}

	for i := 0; i < len(arr); i++ {
		var swap string
		for j := 0; j < d; j++ {
			swap += "\n"
		}

	}
}

func Deep(arr []int) int {
	return int(math.Ceil(math.Log2(float64(len(arr))))) + 1
}
func Test1(t *testing.T) {
	arr := []int{79, 48, 47, 8, 61, 17, 57, 25, 55}
	h := Create(arr)
	fmt.Println(h.Arr)
	Print(h.Arr)
	//h.Insert(80)
	//h.Insert2(70)
	fmt.Println(h.Delete())
	Print(h.Arr)
	fmt.Println(h.Delete())
	Print(h.Arr)
}
