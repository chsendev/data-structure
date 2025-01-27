package data_structure

import (
	"fmt"
	"testing"
)

// 1 1 2 3 5 8
func fib(n, pre2, pre1 int) int {
	if n < 3 {
		return pre1
	}
	return fib(n-1, pre1, pre1+pre2)
}

func TestFib(t *testing.T) {
	fmt.Println(fib(6, 1, 1))
}
