package leetcode_top100

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		j, ok := m[target-num]
		if ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{11, 15, 2, 7}, 9))
}
