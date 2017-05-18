package p460

import (
	"fmt"
)

/**
Design and implement a data structure for Least Frequently Used (LFU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reaches its capacity, it should invalidate the least frequently used item before inserting a new item. For the purpose of this problem, when there is a tie (i.e., two or more keys that have the same frequency), the least recently used key would be evicted.

Follow up:
Could you do both operations in O(1) time complexity?
 */

type LinkMapNode struct {
	val  int
	prev *LinkMapNode
	next *LinkMapNode
}

type DLinkNode struct {
	count int
	keys  map[int]*LinkMapNode
	head  *LinkMapNode
	tail  *LinkMapNode
	prev  *DLinkNode
	next  *DLinkNode
}

func newNode(count int) *DLinkNode {
	h := &LinkMapNode{}
	t := &LinkMapNode{}
	h.next = t
	t.prev = h
	return &DLinkNode{
		count: count,
		keys:  make(map[int]*LinkMapNode),
		head:  h,
		tail:  t}
}

func (n *DLinkNode) addKey(key int) {
	ln := &LinkMapNode{val: key}
	ln.next = n.tail
	ln.prev = n.tail.prev
	n.tail.prev.next = ln
	n.tail.prev = ln
	n.keys[key] = ln
}

func (n *DLinkNode) removeKey(key int) {
	node := n.keys[key]
	node.next.prev = node.prev
	node.prev.next = node.next
	delete(n.keys, key)
}

type LFUCache struct {
	cap     int
	hmNode  map[int]*DLinkNode
	hmValue map[int]int
	head    *DLinkNode
	tail    *DLinkNode
}

func Constructor(capacity int) LFUCache {
	h := &DLinkNode{}
	t := &DLinkNode{}
	h.next = t
	t.prev = h
	return LFUCache{
		cap:     capacity,
		hmNode:  make(map[int]*DLinkNode),
		hmValue: make(map[int]int),
		head:    h,
		tail:    t,
	}
}

func (this *LFUCache) Get(key int) int {
	ans := -1
	if v, ok := this.hmValue[key]; ok {
		this.increaseCount(key)
		ans = v
	}
	return ans
}

func (this *LFUCache) increaseCount(key int) {
	node := this.hmNode[key]
	node.removeKey(key)

	if node.next.count == node.count+1 {
		node.next.addKey(key)
	} else {
		tmp := newNode(node.count + 1)
		tmp.addKey(key)
		tmp.prev = node
		tmp.next = node.next
		node.next = tmp
		tmp.next.prev = tmp
	}
	this.hmNode[key] = node.next
	if len(node.keys) == 0 {
		removeNode(node)
	}
}

func removeNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LFUCache) removeOld() {
	if this.head.next == this.tail {
		return
	}
	old := this.head.next.head.next.val
	this.head.next.removeKey(old)
	if len(this.head.next.keys) == 0 {
		removeNode(this.head.next)
	}
	delete(this.hmValue, old)
	delete(this.hmNode, old)
}

func (this *LFUCache) insertAfterHead(node *DLinkNode) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LFUCache) addAfterHead(key int) {
	if this.head.next == this.tail ||
		this.head.next.count > 0 {
		node := newNode(0)
		node.addKey(key)
		this.insertAfterHead(node)
	} else {
		this.head.next.addKey(key)
	}
	this.hmNode[key] = this.head.next
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	if _, ok := this.hmValue[key]; ok {
		this.hmValue[key] = value
	} else {
		if len(this.hmValue) < this.cap {
			this.hmValue[key] = value
		} else {
			this.removeOld()
			this.hmValue[key] = value
		}
		this.addAfterHead(key)
	}
	this.increaseCount(key)
}

func (this *LFUCache) printList() {
	fmt.Println(this.hmValue)
	for iter := this.head.next; iter != this.tail; iter = iter.next {
		fmt.Print("count ", iter.count, " ", "keys ")
		for cur := iter.head.next; cur != iter.tail; cur = cur.next {
			fmt.Print(cur.val, " ")
		}
		fmt.Print(", ")
	}
	fmt.Println()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
