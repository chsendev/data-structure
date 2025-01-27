package data_structure

import (
	"fmt"
	"math/rand"
	"testing"
)

func GenerateArr() []int {
	arr := make([]int, 11)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
	}
	return arr
}

func TestBubbleSort(t *testing.T) {
	arr := []int{34, 8, 64, 51, 32, 21}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				fmt.Println(arr)
			}
		}
	}
	fmt.Println(arr)
}

func TestSelectSort(t *testing.T) {
	arr := GenerateArr()
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}

func TestInsertSort(t *testing.T) {
	arr := GenerateArr()
	// 1 3 8    4
	for i := 1; i < len(arr); i++ {
		j := i
		temp := arr[i]
		for ; j > 0; j-- {
			if temp < arr[j-1] {
				arr[j] = arr[j-1]
				fmt.Println(arr)
			} else {
				break
			}
		}
		arr[j] = temp
	}
	fmt.Println(arr)
}

func TestShellSort(t *testing.T) {
	arr := GenerateArr()
	fmt.Println(arr)
	//
	for i := len(arr) / 2; i > 0; i /= 2 {
		for j := i; j < len(arr); j++ {
			tmp := arr[j]
			k := j
			for ; k > i; k -= i {
				if tmp < arr[k-i] {
					arr[k] = arr[k-i]
				} else {
					break
				}
			}
			arr[k] = tmp
		}
	}
	fmt.Println(arr)
}

func TestName1(t *testing.T) {
	arr := GenerateArr()

	// [79 48 47 8 61 17 47 17 55 25]
	for d := len(arr) / 2; d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			tmp := arr[i]
			j := i
			for ; j-d >= 0; j -= d {
				if tmp > arr[j-d] {
					arr[j] = arr[j-d]
				} else {
					break
				}
			}
			arr[j] = tmp
		}
	}

	fmt.Println(arr)
	fmt.Printf("\033[1;37;41m%s\033[0m\n", "Red.")
	fmt.Printf("\x1b[38;5;214m这是#f9dd4a颜色的文本\x1b[0m\n")

	fmt.Printf("\033[1;31;40m%s\033[0m\n", "Red.")
	//colorCode := "#f9dd4a"

	// 从十六进制颜色代码转换为RGB
	//r, _ := strconv.ParseInt(strings.TrimPrefix(colorCode[1:3], "0x"), 16, 8)
	//g, _ := strconv.ParseInt(strings.TrimPrefix(colorCode[3:5], "0x"), 16, 8)
	//b, _ := strconv.ParseInt(strings.TrimPrefix(colorCode[5:], "0x"), 16, 8)
	//
	//// 构建ANSI转义序列
	//ansiEscape := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	//// 打印带有颜色的文本
	//text := "这个文本是彩色的"
	//fmt.Println(ansiEscape + text + "\033[0m") // \033[0m 用于重置颜色
	fmt.Printf("\x1b[38;2;%d;%d;%dm这是真彩色模式下的#f9dd4a文本\x1b[0m\n", 249, 221, 74)
}

func CreateMaxHeap(arr []int, n int) {
	// 找到最后一颗子树
	for parent := (n - 1) / 2; parent >= 0; parent-- {
		var child int
		for i := parent; i*2+1 < n; i = child {
			child = i*2 + 1
			if child+1 < n && arr[child] < arr[child+1] {
				child = child + 1
			}
			if arr[child] > arr[parent] {
				arr[parent], arr[child] = arr[child], arr[parent]
			}
		}
	}
}

func TestHeapSort(t *testing.T) {
	arr := GenerateArr()
	for i := 0; i < len(arr)-1; i++ {
		CreateMaxHeap(arr, len(arr)-i)
		arr[len(arr)-i-1], arr[0] = arr[0], arr[len(arr)-i-1]
	}
	fmt.Println(arr)
}

func Merge(arr []int, tmp []int, l, lend, r, rend int) {
	t := l
	size := rend - l + 1
	for l <= lend && r <= rend {
		if arr[l] < arr[r] {
			tmp[t] = arr[l]
			l++
		} else {
			tmp[t] = arr[r]
			r++
		}
		t++
	}
	for l <= lend {
		tmp[t] = arr[l]
		l++
		t++
	}
	for r <= rend {
		tmp[t] = arr[r]
		r++
		t++
	}
	for i := 0; i < size; i++ {
		arr[rend] = tmp[rend]
		rend--
	}
}

func MSort(arr []int, tmp []int, l, r int) {
	if l < r {
		//[79 48 47 8 61 17 47 17 55 25]
		center := (l + r) / 2
		MSort(arr, tmp, l, center)
		MSort(arr, tmp, center+1, r)
		fmt.Printf("正在合并：%d,%d,%d,%d\n", l, center, center+1, r)
		Merge(arr, tmp, l, center, center+1, r)
	}
}

func TestMergeSort(t *testing.T) {
	//arr := GenerateArr()
	arr := []int{45, 43, 30, 19, 21, 53, 0, 37, 88, 82}
	fmt.Println(arr)
	tmp := make([]int, len(arr))
	MSort(arr, tmp, 0, len(arr)-1)
	fmt.Println(arr)
}

func Median3(arr []int, left, right int) int {
	center := (left + right) / 2
	if arr[left] > arr[center] {
		arr[left], arr[center] = arr[center], arr[left]
	}
	if arr[left] > arr[right] {
		arr[left], arr[right] = arr[right], arr[left]
	}
	if arr[center] > arr[right] {
		arr[right], arr[center] = arr[center], arr[right]
	}
	arr[center], arr[right-1] = arr[right-1], arr[center]
	return arr[right-1]
}

func QuickSort(arr []int, left, right int) {
	if right-left <= 1 {
		return
	}

	//pivot := Median3(arr, left, right)
	pivot := arr[right]
	i := left
	j := right - 1
	for {
		for i < right && arr[i] < pivot {
			i++
		}
		for j > left && arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			break
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}

func TestQuickSort(t *testing.T) {
	arr := GenerateArr()
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func TestBucketSort(t *testing.T) {
	//arr := GenerateArr()
	arr := []int{34, 8, 64, 51, 32, 21, 21}
	bucket := make([]*LinkNode, 101)
	for i := 0; i < len(arr); i++ {
		b := bucket[arr[i]]
		n := &LinkNode{Val: arr[i]}
		if b == nil {
			bucket[arr[i]] = n
		} else {
			//bucket[i]
			n.Next = bucket[arr[i]]
			bucket[arr[i]] = n
		}
	}
	for _, node := range bucket {
		if node != nil {
			n := node
			for n != nil {
				fmt.Println(n.Val)
				n = n.Next
			}
		}
	}
}

func TestRadixSort(t *testing.T) {
	arr := GenerateArr() // 生成的位数，必须是3位
	bucket := make([]*LinkNode, 10)
	for i := range bucket {
		bucket[i] = &LinkNode{}
	}
	list := &LinkNode{}
	p := list
	for i := 0; i < len(arr); i++ {
		n := &LinkNode{Val: arr[i]}
		p.Next = n
		p = p.Next
	}
	div := 1
	for i := 0; i < 3; i++ {
		// 取出每一位数
		// 收集
		p = list.Next
		for p != nil {
			num := p.Val / div % 10
			tmp := p
			p = p.Next
			b := bucket[num]
			if b.Next == nil {
				b.Next = tmp
				tmp.Next = nil
			} else {
				for b.Next != nil {
					b = b.Next
				}
				b.Next = tmp
				tmp.Next = nil
				//tmp.Next = b.Next
				//b.Next = tmp
			}
		}
		// 排序
		list.Next = nil
		for j := len(bucket) - 1; j >= 0; j-- {
			if bucket[j] != nil {
				tmp := bucket[j]
				for tmp.Next != nil {
					tmp = tmp.Next
				}
				tmp.Next = list.Next
				list.Next = bucket[j].Next
				bucket[j].Next = nil
			}
		}
		div *= 10
	}
	fmt.Println(bucket)
	p = list.Next
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
