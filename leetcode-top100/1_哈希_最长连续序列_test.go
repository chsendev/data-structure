package leetcode_top100

import (
	"fmt"
	"slices"
	"testing"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]struct{})
	var newNums []int
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			newNums = append(newNums, num)
			m[num] = struct{}{}
		}
	}
	slices.Sort(newNums)

	step := 1
	pre := 1
	for i := 0; i < len(newNums)-1; i++ {
		if newNums[i+1]-newNums[i] == 1 {
			pre += 1
		} else if step < pre {
			step = pre
			pre = 1
		} else {
			pre = 1
		}
	}
	if step < pre {
		step = pre
		pre = 1
	}
	return step
}

func TestLongestConsecutive(t *testing.T) {
	fmt.Println(longestConsecutive([]int{9, 1, -3, 2, 4, 8, 3, -1, 6, -2, -4, 7}))
}
