package leetcode_top100

import (
	"fmt"
	"testing"
)

var prime2 = []int{5, 71, 31, 29, 2, 53, 59, 23, 11, 89, 79, 37, 41, 13, 7, 43, 97, 17, 19, 3, 47, 73, 61, 83, 67, 101}

func findAnagrams(s string, p string) []int {
	sum := 1
	for i := 0; i < len(p); i++ {
		sum *= prime2[int(p[i])-97]
	}

	var left int
	var res []int
	for i := 0; i < len(s); i++ {
		if i-left < len(p)-1 {
			continue
		}

		t := 1
		for j := left; j <= i; j++ {
			t *= prime2[int(s[j])-97]
		}
		if t == sum {
			res = append(res, left)
		} else {
			left = i
		}
		left++
	}

	return res
}

func findAnagrams2(s string, p string) []int {
	need := make(map[int]int)
	window := make(map[int]int)
	for i := 0; i < len(p); i++ {
		need[int(p[i])]++
	}
	var left, right int
	var res []int
	var valid int
	for right < len(s) {
		c := int(s[right])
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		right++

		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			d := int(s[left])
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return res
}

func TestFindAnagrams(t *testing.T) {
	// c b d a e b a
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	//fmt.Println(findAnagrams2("cdbaebabacd", "abc"))
}
