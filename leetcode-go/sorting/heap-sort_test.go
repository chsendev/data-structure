package sorting

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/sort-an-array/

func TestSortArray3(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1, 6}
	fmt.Println(sortArray3(nums))
}

func TestSortArray4(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1, 6}
	fmt.Println(sortArray4(nums))
}

// 从底到顶
func sortArray4(nums []int) []int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, i, len(nums))
		//heapInsert(nums, i)
	}
	size := len(nums)
	for size > 1 {
		nums[0], nums[size-1] = nums[size-1], nums[0]
		size--
		heapify(nums, 0, size)
	}
	return nums
}

// 从顶到底
func sortArray3(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		heapInsert(nums, i)
	}
	size := len(nums)
	for size > 1 {
		nums[0], nums[size-1] = nums[size-1], nums[0]
		size--
		heapify(nums, 0, size)
	}
	return nums
}

func TestHeapInsert(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1, 6}
	heapInsert(nums, 5)
	fmt.Println(nums)
}

// 大顶堆，向上调整
func heapInsert(nums []int, i int) {
	h := (i - 1) / 2
	for nums[i] > nums[h] {
		nums[h], nums[i] = nums[i], nums[h]
		i = h
		h = (i - 1) / 2
	}
}

/*
		4
	2		3

0		1
*/
func TestHeapify(t *testing.T) {
	nums := []int{0, 4, 3, 2, 1}
	heapify(nums, 0, 5)
	fmt.Println(nums)
}

// 大顶堆，向下调整
func heapify(nums []int, i int, size int) {
	l := 2*i + 1
	for l < size {
		best := l
		r := 2*i + 2
		if r < size && nums[r] > nums[l] {
			best = r
		}
		if nums[best] <= nums[i] {
			break
		}
		nums[best], nums[i] = nums[i], nums[best]
		i = best
		l = 2*i + 1
	}
}
