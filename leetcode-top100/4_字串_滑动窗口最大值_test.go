package leetcode_top100

import (
	"fmt"
	"math"
	"testing"
)

func maxSlidingWindow(nums []int, k int) []int {
	maxNum := math.MinInt
	maxIndex := 0
	for i := k - 1; i >= 0; i-- {
		if maxNum < nums[i] {
			maxNum = nums[i]
			maxIndex = i
		}
	}
	res := []int{maxNum}
	for i := k; i < len(nums); i++ {
		if nums[i] >= maxNum {
			// 比之前最大值还大
			maxNum = nums[i]
			maxIndex = i
			res = append(res, maxNum)
		} else {
			// 比之前最大值小
			if i-maxIndex < k {
				// 最大值没超过区间
				res = append(res, maxNum)
			} else {
				maxNum = math.MinInt
				t := maxIndex + 1
				for j := i; j >= t; j-- {
					if maxNum < nums[j] {
						maxNum = nums[j]
						maxIndex = j
					}
				}
				res = append(res, maxNum)
			}
		}
	}
	return res
}

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}
