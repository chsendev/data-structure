package leetcode_top100

import (
	"fmt"
	"testing"
)

func productExceptSelf(nums []int) []int {
	pre, post := 1, 1
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = pre
		pre *= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= post
		post *= nums[i]
	}
	return res
}

func TestProductExceptSelf(t *testing.T) {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
