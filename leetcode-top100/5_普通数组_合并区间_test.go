package leetcode_top100

import (
	"fmt"
	"slices"
	"testing"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a []int, b []int) int {
		return a[0] - b[0]
	})
	fmt.Println(intervals)

	res := [][]int{{intervals[0][0], intervals[0][1]}}
	idx := 0
	for i := 1; i < len(intervals); i++ {
		if res[idx][1] >= intervals[i][0] && res[idx][0] <= intervals[i][1] {
			if res[idx][0] > intervals[i][0] {
				res[idx][0] = intervals[i][0]
			}
			if res[idx][1] < intervals[i][1] {
				res[idx][1] = intervals[i][1]
			}
		} else {
			res = append(res, []int{intervals[i][0], intervals[i][1]})
			idx++
		}
	}
	return res
}

func TestMerge(t *testing.T) {
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}
