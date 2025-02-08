package data_structure

import (
	"fmt"
	"testing"
)

// 整型数组，把偶数调整到数组的左边，把奇数调整到数组的右边

func AdjustArray(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		for l < r && arr[l]%2 == 0 {
			// 偶数
			l++
		}
		for l < r && arr[r]%2 == 1 {
			// 奇数
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
}

func TestAdjustArray(t *testing.T) {
	arr := GenerateArr()
	fmt.Println(arr)
	AdjustArray(arr)
	fmt.Println(arr)
}
