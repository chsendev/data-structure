package leetcode_top100

import (
	"fmt"
	"testing"
)

func moveZeroes(nums []int) {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			t := nums[i]
			nums[i] = 0
			nums[n] = t
			n++
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
