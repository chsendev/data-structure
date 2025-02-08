package leetcode_top100

import (
	"fmt"
	"testing"
)

var prime = []int{5, 71, 31, 29, 2, 53, 59, 23, 11, 89, 79, 37, 41, 13, 7, 43, 97, 17, 19, 3, 47, 73, 61, 83, 67, 101}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	m := make(map[int][]string)
	for _, s1 := range strs {
		sum := 1
		for _, s2 := range s1 {
			sum *= prime[int(s2)-97]
		}
		m[sum] = append(m[sum], s1)
	}

	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func TestGroupAnagrams(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
