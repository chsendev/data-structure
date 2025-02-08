package data_structure

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type CircularQueue struct {
	data  []int
	front atomic.Int64
	rear  atomic.Int64
	sync.Mutex
}

func NewCircularQueue(length int) *CircularQueue {
	return &CircularQueue{
		data:  make([]int, length),
		front: atomic.Int64{},
		rear:  atomic.Int64{},
	}
}

func (c *CircularQueue) IsEmpty() bool {
	return c.front == c.rear
}

func (c *CircularQueue) IsFull() bool {
	// 1 2 3 4 5
	//   r f
	//
	r := (c.rear.Load() + 1) % int64(len(c.data))
	return c.front.Load() == r
}

func (c *CircularQueue) Push(d int) error {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.IntN(100)))
		nextRear := (c.rear.Load() + 1) % int64(len(c.data))
		if c.front.Load() == nextRear {
			return errors.New("is full")
		}
		if c.rear.CompareAndSwap(c.rear.Load(), nextRear) {
			c.data[c.rear.Load()] = d
			return nil
		}
	}
}

func (c *CircularQueue) Pop() (int, error) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.IntN(100)))
		if c.front.Load() == c.rear.Load() {
			return 0, errors.New("is empty")
		}
		f := c.front.Load()
		nextFront := (c.front.Load() + 1) % int64(len(c.data))
		if c.front.CompareAndSwap(c.front.Load(), nextFront) {
			return c.data[f], nil
		}
	}
}

func TestCircularQueue(t *testing.T) {
	var l atomic.Int64
	var l2 atomic.Int64
	l.Add(1)
	l2.Add(1)
	fmt.Println(l.Load() == l2.Load())
	l.Add(1)
	l.Add(1)
	fmt.Println(l.Load())

	cq := NewCircularQueue(10 * 3)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(cq.Push(1))
			fmt.Println(cq.Push(2))
			fmt.Println(cq.Push(3))
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			//fmt.Println(cq.Pop())
			//fmt.Println(cq.Pop())
			//fmt.Println(cq.Pop())
			//fmt.Println(cq.Pop())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(len(cq.data))
	//fmt.Println((cq.rear - cq.front + 1 + len(cq.data)) % len(cq.data))
}
