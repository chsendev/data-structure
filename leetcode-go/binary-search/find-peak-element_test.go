package binary_search

import (
	"fmt"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	nums := []int{1, 3, 2, 1}
	fmt.Println(findPeakElement(nums))
}

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	l := 1
	r := len(nums) - 1
	m := 0
	ans := -1
	for l <= r {
		m = l + (r-l)/2
		if nums[m-1] > nums[m] {
			r = m - 1
		} else if nums[m] < nums[m+1] {
			l = m + 1
		} else {
			ans = m
			break
		}
	}
	return ans
}
