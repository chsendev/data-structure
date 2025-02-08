package leetcode_top100

import (
	"fmt"
	"testing"
)

func setZeroes2(matrix [][]int) {
	cols := make(map[int]struct{})
	rows := make(map[int]struct{})
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				_, ok := cols[j]
				if !ok {
					cols[j] = struct{}{}
				}
				_, ok = rows[i]
				if !ok {
					rows[i] = struct{}{}
				}
			}
		}
	}
	for k := range rows {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[k][j] = 0
		}
	}
	for k := range cols {
		for j := 0; j < len(matrix); j++ {
			matrix[j][k] = 0
		}
	}
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	var zeros [][]int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				zeros = append(zeros, []int{i, j})
			}
		}
	}
	for i := 0; i < len(zeros); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[j][zeros[i][1]] = 0
		}
		for j := 0; j < len(matrix[0]); j++ {
			matrix[zeros[i][0]][j] = 0
		}
	}
	fmt.Println(matrix)
}

func TestSetZeroes(t *testing.T) {
	setZeroes([][]int{
		{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5},
	})
}
