package data_structure

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync"
	"testing"
	"time"
)

type MaxHeap struct {
	Arr []int
}

func Create(arr []int) *MaxHeap {
	var h MaxHeap
	// 找到最后一颗子树
	for parent := (len(arr) - 1) / 2; parent >= 0; parent-- {
		var child int
		for i := parent; i*2+1 < len(arr); i = child {
			child = i*2 + 1
			if child+1 < len(arr) && arr[child] < arr[child+1] {
				child = child + 1
			}
			if arr[child] > arr[parent] {
				arr[parent], arr[child] = arr[child], arr[parent]
			}
		}
	}
	h.Arr = arr
	return &h
}

func (m *MaxHeap) Insert(item int) {
	m.Arr = append(m.Arr, item)
	var parent int
	for child := len(m.Arr) - 1; (child-1)/2 >= 0; child = parent {
		// 找到父节点
		parent = (child - 1) / 2
		// 是否比父节点大，是的话则交换
		if m.Arr[child] > m.Arr[parent] {
			m.Arr[child], m.Arr[parent] = m.Arr[parent], m.Arr[child]
		} else {
			break
		}
	}
}

func (m *MaxHeap) Insert2(item int) {
	m.Arr = append(m.Arr, item)
	m.Recursion(len(m.Arr) - 1)

}
func (m *MaxHeap) Recursion(child int) {
	if (child-1)/2 < 0 {
		return
	}
	parent := (child - 1) / 2
	if m.Arr[child] > m.Arr[parent] {
		m.Arr[child], m.Arr[parent] = m.Arr[parent], m.Arr[child]
		m.Recursion(parent)
	} else {
		return
	}
}

func (m *MaxHeap) Delete() int {
	maxEle := m.Arr[0]
	m.Arr[0] = m.Arr[len(m.Arr)-1]
	m.Arr = m.Arr[:len(m.Arr)-1]
	var child int
	for parent := 0; parent*2+1 < len(m.Arr); parent = child {
		child = parent*2 + 1
		if parent*2+2 < len(m.Arr) && m.Arr[child] < m.Arr[parent*2+2] {
			child = parent*2 + 2
		}
		if m.Arr[child] > m.Arr[parent] {
			m.Arr[child], m.Arr[parent] = m.Arr[parent], m.Arr[child]
		} else {
			break
		}
	}
	return maxEle
}

func Print(arr []int) {
	d := Deep(arr)
	for i := 0; i < d; i++ {
		layer := int(math.Pow(2, float64(i)))
		first := int(math.Pow(2, float64(i))) - 1
		var swap string
		for j := i; j < d; j++ {
			swap += "\t"
		}
		fmt.Print(swap)
		for j := first; j < layer+first && j < len(arr); j++ {
			fmt.Printf(strconv.Itoa(int(arr[j])))
			swap = ""
			for k := i; k < d; k++ {
				swap += "  "
			}
			fmt.Print(swap)
		}
		fmt.Println()
	}

	for i := 0; i < len(arr); i++ {
		var swap string
		for j := 0; j < d; j++ {
			swap += "\n"
		}

	}
}

func Deep(arr []int) int {
	return int(math.Ceil(math.Log2(float64(len(arr))))) + 1
}

func TestGo(t *testing.T) {
	//ch := make(chan int, 10)
	l := sync.Mutex{}

	cond := sync.NewCond(&l)
	cond.L.Lock()

	go func() {
		cond.Wait()
		fmt.Println("被唤醒")

		//for {
		//	select {
		//	case <-ch:
		//		fmt.Println(1111)
		//	}
		//}
	}()
	go func() {
		cond.Wait()
		fmt.Println("被唤醒")

	}()

	time.Sleep(time.Second)

	cond.Broadcast()

	time.Sleep(time.Hour)
}

func Test1(t *testing.T) {
	(&http.Client{Timeout: time.Second * 2250}).Get("https://res.jogg.ai/joggUserData/project/ae7f59bd6002476b94de14d5dc16027b/1738996833059-359c7d0a8f6cc89ce3b4480dd030aeda-video.mp4")
	time.Sleep(time.Hour)
	arr := []int{79, 48, 47, 8, 61, 17, 57, 25, 55}
	h := Create(arr)
	fmt.Println(h.Arr)
	Print(h.Arr)
	//h.Insert(80)
	//h.Insert2(70)
	fmt.Println(h.Delete())
	Print(h.Arr)
	fmt.Println(h.Delete())
	Print(h.Arr)
}

func TestTh(t *testing.T) {
	i := 0
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go func() {
		for j := 0; j < 10; j++ {
			<-ch1
			i++
			fmt.Println(i)
			ch2 <- struct{}{}
		}
	}()
	go func() {
		for j := 0; j < 10; j++ {
			<-ch2
			i--
			fmt.Println(i)
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}

	time.Sleep(time.Hour)
}

func TestTh2(t *testing.T) {
	i := 0
	b := false
	cond := sync.NewCond(&sync.RWMutex{})
	go func() {
		for j := 0; j < 10; j++ {
			cond.L.Lock()
			for b {
				fmt.Println("协诚A等待被唤醒")
				cond.Wait()
				fmt.Println("协诚A被唤醒")
			}
			i++
			fmt.Println(i)
			b = true
			cond.Broadcast()
			fmt.Println("协诚A主动唤醒别人")
			cond.L.Unlock()
		}
	}()
	go func() {
		for j := 0; j < 10; j++ {
			cond.L.Lock()
			for !b {
				fmt.Println("协诚B等待被唤醒")
				cond.Wait()
				fmt.Println("协诚B被唤醒")
			}
			i--
			fmt.Println(i)
			b = false
			cond.Broadcast()
			fmt.Println("协诚B主动唤醒别人")
			cond.L.Unlock()
		}
	}()

	time.Sleep(time.Hour)
}

//type Node struct {
//	Data int
//}

func TestPriorityQueue(t *testing.T) {
	pq := PriorityQueue{
		Data: make([]int, 0),
		cond: sync.NewCond(&sync.Mutex{}),
	}
	go func() {
		for {
			fmt.Println("协程1", pq.Pop())
		}
	}()
	go func() {
		for {
			fmt.Println("协程2", pq.Pop())
		}
	}()
	go func() {
		for {
			fmt.Println("协程3", pq.Pop())
		}
	}()
	time.Sleep(time.Second)
	pq.Push(1)
	time.Sleep(time.Second)
	pq.Push(2)
	time.Sleep(time.Second)
	pq.Push(3)
	time.Sleep(time.Second)
	pq.Push(4)
	time.Sleep(time.Second)
	pq.Push(5)
	time.Sleep(time.Second)
	pq.Push(6)
	time.Sleep(time.Second)
	pq.Push(7)
	time.Sleep(time.Second)
	pq.Push(8)

	time.Sleep(time.Hour)
	//fmt.Println(pq.Data)
	//fmt.Println(pq.Pop())
	//fmt.Println(pq.Data)
}

type PriorityQueue struct {
	Data []int
	cond *sync.Cond
}

func (p *PriorityQueue) Push(d int) {
	p.heapInsert(d)
	p.cond.Broadcast()
}

func (p *PriorityQueue) Pop() int {
	p.cond.L.Lock()
	for len(p.Data) == 0 {
		p.cond.Wait()
	}
	defer p.cond.L.Unlock()

	d := p.Data[0]
	p.Data[0] = p.Data[len(p.Data)-1]
	p.Data = p.Data[:len(p.Data)-1]
	p.heapify(0)
	return d
}

func (p *PriorityQueue) heapInsert(d int) {
	p.Data = append(p.Data, d)
	idx := len(p.Data) - 1
	parent := (idx - 1) / 2
	for parent >= 0 {
		if p.Data[idx] > p.Data[parent] {
			p.Data[idx], p.Data[parent] = p.Data[parent], p.Data[idx]
			idx = parent
			parent = (idx - 1) / 2
		} else {
			break
		}
	}
}

func (p *PriorityQueue) heapify(i int) {
	l := i*2 + 1
	for l < len(p.Data) {
		best := l
		r := l + 1
		if r < len(p.Data) && p.Data[r] > p.Data[l] {
			best = r
		}
		if p.Data[i] > p.Data[best] {
			break
		}
		p.Data[i], p.Data[best] = p.Data[best], p.Data[i]
		i = best
		l = i*2 + 1
	}
}

func TestCh(t *testing.T) {
	//var ch chan int = nil
	ch := make(chan int, 5)

	close(ch)
	close(ch)

	u := struct {
		ch chan int
	}{}
	u.ch = make(chan int, 5)

	go func() {
		//var ch chan int = nil
		//ch <- 10
		//<-ch

		u.ch <- 1
		u.ch <- 2
		u.ch <- 3
		u.ch <- 4
		u.ch <- 5
		u.ch <- 6

		//for c := range ch {
		//	fmt.Println(c)
		//}
		fmt.Println("ok")
	}()
	go func() {
		time.Sleep(time.Second)
		//close(u.ch)
		<-u.ch
		u.ch = nil
	}()
	time.Sleep(time.Hour)
}

func TestFor(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	tmp := 0
	for i := 0; i < 100; i++ {
		go func(y int) {
			cond.L.Lock()
			for tmp != y {
				cond.Wait()
			}
			fmt.Println(tmp)
			tmp++
			cond.Broadcast()
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(time.Hour)

	//wg := sync.WaitGroup{}
	//wg.Add(10)
	//chs := make(chan int)
	////通过一个全局变量控制进channel的顺序
	//tag := 1
	//for i := 1; i <= 10; i++ {
	//	go func(value int) {
	//		//死循环，保证按顺序进chan
	//		for {
	//			if tag == value {
	//				chs <- value
	//				break
	//			}
	//		}
	//	}(i)
	//}
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-chs)
	//	wg.Done()
	//	tag++
	//}
	//wg.Wait()

	////cond := sync.NewCond(&sync.Mutex{})
	//ch := make(chan int, 1)
	//for i := 0; i < 100; i++ {
	//	//x := i
	//	ch <- i
	//	go func(y int) {
	//		fmt.Println(<-ch)
	//		//for x != y {
	//		//	cond.Wait()
	//		//}
	//		//fmt.Println(x, y)
	//		//cond.Broadcast()
	//	}(i)
	//}
	//
	//time.Sleep(time.Hour)
}
