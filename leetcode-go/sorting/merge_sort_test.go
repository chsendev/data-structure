package sorting

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

var help []int

func TestMergeSort(t *testing.T) {
	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	help = make([]int, len(nums))
	//mergeSort1(nums, 0, len(nums)-1)
	mergeSort2(nums)
	return nums
}

// [46 20 65 38 87 17 48 94 45 5]
//
//	0  1  2  3  4   5  6  7  8  9
//
// mergeSort2 非递归版
func mergeSort2(nums []int) {
	var l, m, r int
	for step := 1; step < len(nums); step <<= 1 {
		l = 0
		fmt.Println("step = ", step)
		for l < len(nums) {
			m = l + step - 1
			if (m + 1) >= len(nums) {
				break
			}
			r = int(math.Min(float64(len(nums)-1), float64(m+step)))
			fmt.Println("l =", l, ",m = ", m, ",r = ", r)
			merge(nums, l, m, r)
			l = r + 1
		}
	}
}

// mergeSort1 递归版
func mergeSort1(nums []int, l, r int) {
	if r == l {
		return
	}
	m := (r + l) / 2
	mergeSort1(nums, l, m)   // 合并左边
	mergeSort1(nums, m+1, r) // 合并右边
	merge(nums, l, m, r)
}

func merge(nums []int, l, m, r int) {
	x, y := l, m+1 // 左右指针
	i := l
	for ; x <= m && y <= r; i++ {
		if nums[x] <= nums[y] {
			help[i] = nums[x]
			x++
		} else {
			help[i] = nums[y]
			y++
		}
	}
	for ; x <= m; i++ {
		help[i] = nums[x]
		x++
	}
	for ; y <= r; i++ {
		help[i] = nums[y]
		y++
	}
	for i = l; i <= r; i++ {
		nums[i] = help[i]
	}
}
