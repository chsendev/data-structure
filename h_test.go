package data_structure

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	for i := 1; i < len(arr); i++ {
		heapInsert(arr, i)
	}

	size := len(arr)
	for size > 1 {
		arr[0], arr[size-1] = arr[size-1], arr[0]
		size--
		heapify2(arr, 0, size)
	}

	fmt.Println(arr)
}

func heapInsert(arr []int, i int) []int {
	parent := (i - 1) / 2
	for parent >= 0 {
		if arr[parent] < arr[i] {
			arr[parent], arr[i] = arr[i], arr[parent]
			i = parent
			parent = (i - 1) / 2
		} else {
			break
		}
	}
	return arr
}

func heapify2(arr []int, i, size int) {
	left := 2*i + 1

	for left < size {
		best := left
		right := left + 1
		if right < size && arr[left] < arr[right] {
			best = right
		}

		if arr[i] < arr[best] {
			arr[i], arr[best] = arr[best], arr[i]
			i = left
			left = (i + 1) * 2
		} else {
			break
		}
	}
}

func part(arr []int, left, right, x int) int {

	xi := 0
	a := 0
	for i := left; i <= right; i++ {
		if arr[i] <= x {
			arr[i], arr[a] = arr[a], arr[i]
			if arr[a] == x {
				xi = a
			}
			a++
		}
	}

	arr[a-1], arr[xi] = arr[xi], arr[a-1]
	return a - 1
}

func TestPart(t *testing.T) {
	arr := []int{3, 4, 6, 5, 2, 8, 1}
	part(arr, 0, len(arr)-1, 5)
	fmt.Println(arr)
}

func TestQuickSort1(t *testing.T) {
	arr := []int{3, 4, 6, 5, 2, 8, 1}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	x := arr[(left+right)/2]

	m := partition(arr, left, right, x)

	quickSort(arr, left, m-1)
	quickSort(arr, m+1, right)
}

func partition(arr []int, l, r, x int) int {
	a := l
	xi := 0
	for i := l; i <= r; i++ {
		if arr[i] <= x {
			arr[a], arr[i] = arr[i], arr[a]
			if arr[a] == x {
				xi = a
			}
			a++
		}
	}
	arr[xi], arr[a-1] = arr[a-1], arr[xi]
	return a - 1
}
