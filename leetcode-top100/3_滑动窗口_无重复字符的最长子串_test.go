package leetcode_top100

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	var left, right, maxLength int
	m := make(map[int]struct{})
	for ; right < len(s); right++ {
		for _, ok := m[int(s[right])]; ok; _, ok = m[int(s[right])] {
			delete(m, int(s[left]))
			left++
		}
		m[int(s[right])] = struct{}{}
		if (right - left + 1) > maxLength {
			maxLength = right - left + 1
		}
	}
	return maxLength
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abccabcbb"))
}
