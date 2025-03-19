package leetcode_top100

import (
	"container/list"
	"fmt"
	"testing"
)

type pair struct {
	key   int
	value any
}

type LRUCache struct {
	cache    *list.List
	hash     map[int]*list.Element
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    list.New(),
		hash:     make(map[int]*list.Element),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	ele, ok := this.hash[key] // 查找key对应的节点
	if !ok {
		return -1
	}
	this.cache.Remove(ele) // 删除该节点
	value := ele.Value.(pair).value.(int)
	this.hash[key] = this.cache.PushFront(pair{key: key, value: value}) // 重新插入头部，并更新哈希
	return value
}

func (this *LRUCache) Put(key int, value int) {
	ele, ok := this.hash[key]
	if ok { // 缓存中有key
		this.cache.Remove(ele)
	} else if this.cache.Len() >= this.capacity { // 缓存中没有key，但容量不够
		old := this.cache.Back()
		this.cache.Remove(old)
		delete(this.hash, old.Value.(pair).key)
	}
	this.hash[key] = this.cache.PushFront(pair{key: key, value: value}) // 将key对应的节点插入队列头部，更新哈希
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func TestLRUCache(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(2, 1)
	lruCache.Put(1, 1)
	fmt.Println(lruCache.Get(2))
	lruCache.Put(4, 1)
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(2))
}
