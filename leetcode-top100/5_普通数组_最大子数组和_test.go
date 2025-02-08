package leetcode_top100

import (
	"fmt"
	"math"
	"testing"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	sum := 0
	res := math.MinInt

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > res {
			res = sum
		}
		if sum <= 0 {
			sum = 0
		}
	}
	return res
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, -1}))
}
