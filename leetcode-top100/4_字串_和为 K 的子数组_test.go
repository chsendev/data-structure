package leetcode_top100

import (
	"fmt"
	"testing"
)

// 前缀和 https://www.bilibili.com/video/BV1gN411E7Zx/?spm_id_from=333.337.search-card.all.click&vd_source=692f881591c8b7932baa8203eda07c2d

func subarraySum(nums []int, k int) int {
	m := map[int]int{
		0: 1,
	}
	pre := 0
	res := 0
	for _, v := range nums {
		pre += v

		v2, ok := m[pre-k]
		if ok {
			res += v2
		}
		m[pre]++
	}
	return res
}

func subarraySum2(nums []int, k int) int {
	m := make(map[int]int)
	t := 0
	for _, v := range nums {
		t += v
		m[t]++
	}
	res := 0
	for key := range m {
		if key == k {
			res++
		}
		_, ok := m[key-k]
		if ok {
			res++
		}
	}
	return res
}

func TestSubarraySum(t *testing.T) {
	fmt.Println(subarraySum([]int{1}, 0))
}
