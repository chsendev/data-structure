package sorting

import (
	"fmt"
	"math/rand"
	"testing"
)

// https://leetcode.cn/problems/kth-largest-element-in-an-array/

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 2))
}

func findKthLargest(nums []int, k int) int {
	var ans int
	find(nums, 0, len(nums)-1, len(nums)-k, &ans)
	return ans
}

func find(nums []int, l, r, k int, ans *int) {
	if l == r {
		*ans = nums[l]
		return
	} else if l > r {
		return
	}
	i, j := partition2(nums, l, r)
	if k >= i && k <= j {
		*ans = nums[k]
	} else if k < i {
		find(nums, l, i-1, k, ans)
	} else {
		find(nums, j+1, r, k, ans)
	}
}

func TestPartition2(t *testing.T) {
	nums := []int{1, 3, 5, 2, 8, 6, 5}
	fmt.Println(partition2(nums, 0, len(nums)-1))
}

func partition2(nums []int, l, r int) (int, int) {
	x := nums[l+rand.Intn(r-l+1)]
	i, j := l, r
	k := l
	for k <= j {
		if nums[k] < x {
			nums[k], nums[i] = nums[i], nums[k]
			i++
			k++
		} else if nums[k] > x {
			nums[k], nums[j] = nums[j], nums[k]
			j--
		} else {
			k++
		}
	}
	return i, j
}
