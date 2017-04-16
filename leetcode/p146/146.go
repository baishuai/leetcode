package p146

import "fmt"

/**
Design and implement a data structure for Least Recently Used (LRU) cache.
 It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key
 exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present.
 When the cache reached its capacity, it should invalidate the least
 recently used item before inserting a new item.

Follow up:
Could you do both operations in O(1) time complexity?
 */

type DLinkNode struct {
	key  int
	val  int
	prev *DLinkNode
	next *DLinkNode
}

type LRUCache struct {
	cap  int
	len  int
	hm   map[int]*DLinkNode
	head *DLinkNode
	tail *DLinkNode
}

func Constructor(capacity int) LRUCache {
	h := &DLinkNode{}
	t := &DLinkNode{}
	h.next = t
	t.prev = h
	return LRUCache{
		cap:  capacity,
		len:  0,
		hm:   make(map[int]*DLinkNode),
		head: h,
		tail: t,
	}
}

func (this *LRUCache) Get(key int) int {
	ans := -1
	if v, ok := this.hm[key]; ok {
		ans = v.val
		v.prev.next = v.next
		v.next.prev = v.prev
		this.insertAfterHead(v)
	}
	this.printList()
	return ans
}

func (this *LRUCache) insertAfterHead(node *DLinkNode) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) printList() {
	for iter := this.head.next; iter != this.tail; iter = iter.next {
		fmt.Printf("(%d,%d) ", iter.key, iter.val)
	}
	fmt.Println()
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.hm[key]; ok {
		v.val = value
		v.prev.next = v.next
		v.next.prev = v.prev
		this.insertAfterHead(v)
	} else if this.len < this.cap {
		node := &DLinkNode{key: key, val: value}
		this.insertAfterHead(node)
		this.hm[key] = node
		this.len++
	} else {
		node := this.tail.prev
		delete(this.hm, node.key)
		node.key = key
		node.val = value
		node.prev.next = node.next
		node.next.prev = node.prev
		this.insertAfterHead(node)
		this.hm[key] = node
	}
	this.printList()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
