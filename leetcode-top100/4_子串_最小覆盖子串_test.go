package leetcode_top100

import (
	"fmt"
	"testing"
)

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	m := make(map[int]int)
	for _, v := range t {
		m[int(v)]++
	}

	m2 := make(map[int]int)

	minSub := ""
	var valid int
	var left int
	for i := 0; i < len(s); i++ {
		c := int(s[i])
		if _, ok := m[c]; ok {
			m2[c]++
			if m[c] == m2[c] {
				valid++
			}
		}
		for valid == len(m) {
			for _, ok := m[int(s[left])]; !ok; _, ok = m[int(s[left])] {
				left++
			}
			sub := s[left : i+1]
			if minSub == "" || len(sub) < len(minSub) {
				minSub = sub
			}

			if m[int(s[left])] == m2[int(s[left])] {
				valid--
			}
			m2[int(s[left])]--
			left++
		}
	}
	return minSub
}

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("a", "a"))
}
