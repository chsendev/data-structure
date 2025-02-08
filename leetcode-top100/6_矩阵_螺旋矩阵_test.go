package leetcode_top100

import (
	"fmt"
	"testing"
)

func spiralOrder(matrix [][]int) []int {
	sum := len(matrix) * len(matrix[0])
	res := make([]int, 0, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i])-i; j++ {
			res = append(res, matrix[i][j])
		}
		if len(res) == sum {
			return res
		}
		for j := i + 1; j < len(matrix[0])-i-1; j++ {
			res = append(res, matrix[j][len(matrix)-1-i])
		}
		if len(res) == sum {
			return res
		}
		for j := len(matrix) - 1 - i; j > i-1; j-- {
			res = append(res, matrix[len(matrix)-1-i][j])
		}
		if len(res) == sum {
			return res
		}
		for j := len(matrix) - i - 1 - 1; j > 0+i; j-- {
			res = append(res, matrix[j][i])
		}
		if len(res) == sum {
			return res
		}
	}
	return res
}

func TestSpiralOrder(t *testing.T) {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		//{11, 12, 13, 14, 15},
		//{16, 17, 18, 19, 20},
		//{21, 22, 23, 24, 25},
		//{26, 27, 28, 29, 30},
		//{31, 32, 33, 34, 35},
	}))
}
