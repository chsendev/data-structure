package leetcode_top100

import (
	"fmt"
	"math"
	"testing"
)

func maxArea(height []int) int {
	res := 0
	l, r := 0, len(height)-1
	for l < r {
		if height[l] < height[r] {
			t := height[l] * (r - l)
			if res < t {
				res = t
			}
			l++
		} else {
			t := height[r] * (r - l)
			if res < t {
				res = t
			}
			r--
		}
	}
	return res
}

func maxArea2(height []int) int {
	res := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			t := int(math.Min(float64(height[i]), float64(height[j]))) * (j - i)
			if res < t {
				res = t
			}
		}
	}
	return res
}

func TestMaxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
