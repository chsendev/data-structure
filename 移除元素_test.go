package data_structure

import (
	"fmt"
	"testing"
)

func removeElement(nums []int, val int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		for l <= r && nums[r] == val {
			r--
		}
		for l <= r && nums[l] != val {
			l++
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return l
}

func TestRemoveElement(t *testing.T) {
	fmt.Println(removeElement([]int{2}, 3))
}
