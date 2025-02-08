package leetcode_top100

import (
	"fmt"
	"slices"
	"testing"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	slices.Sort(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && i > 0 && nums[i-1] == nums[i] {
			continue
		}

		for j := i + 1; j < len(nums)-1; j++ {
			if i == j || (j > i+1 && nums[j-1] == nums[j]) {
				continue
			}

			for k := j - 1; k >= 0; k-- {
				if i == k || j == k || (k < j-1 && nums[k+1] == nums[k]) {
					continue
				}

				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return res
}

func TestThreeSum(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
