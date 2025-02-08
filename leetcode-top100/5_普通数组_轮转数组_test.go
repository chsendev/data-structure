package leetcode_top100

import (
	"testing"
)

func rotate(nums []int, k int) {
	k %= len(nums)
	temp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		temp[i] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = temp[(len(nums)-k+i)%len(nums)]
	}
}

func TestRotate(t *testing.T) {
	rotate([]int{1, 2}, 3)
	// 6 7 4 5 1 2 3
}
