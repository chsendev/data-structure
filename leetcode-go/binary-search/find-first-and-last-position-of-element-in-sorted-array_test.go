package binary_search

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	if len(nums) == 0 {
		return ans
	}
	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}

	l := 0
	r := len(nums) - 1
	m := 0
	for l <= r {
		m = (l + r) / 2
		if nums[m] == target {
			r = m - 1
			ans[0] = m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	l = 0
	r = len(nums) - 1
	m = 0
	for l <= r {
		m = (l + r) / 2
		if nums[m] == target {
			l = m + 1
			ans[1] = m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans
}
