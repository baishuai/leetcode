package p432

import "math"

/**
Implement a data structure supporting the following operations:

Inc(Key) - Inserts a new key with value 1. Or increments an existing key by 1. Key is guaranteed to be a non-empty string.
Dec(Key) - If Key's value is 1, remove it from the data structure. Otherwise decrements an existing key by 1. If the key does not exist, this function does nothing. Key is guaranteed to be a non-empty string.
GetMaxKey() - Returns one of the keys with maximal value. If no element exists, return an empty string "".
GetMinKey() - Returns one of the keys with minimal value. If no element exists, return an empty string "".
Challenge: Perform all these in O(1) time complexity.
 */

type DlinkNode struct {
	key  string
	val  int
	prev *DlinkNode
	next *DlinkNode
}

type AllOne struct {
	hm   map[string]*DlinkNode
	head *DlinkNode
	tail *DlinkNode
}

func swapAdjacent(p, n *DlinkNode) {
	p.prev.next = n
	n.next.prev = p
	n.prev = p.prev
	p.next = n.next
	n.next = p
	p.prev = n
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	h := &DlinkNode{val: math.MaxInt32}
	t := &DlinkNode{val: 0}
	h.next = t
	t.prev = h

	return AllOne{
		hm:   make(map[string]*DlinkNode),
		head: h,
		tail: t,
	}
}

func (this*AllOne) insertBeforeTail(node *DlinkNode) {
	node.next = this.tail
	node.prev = this.tail.prev
	node.prev.next = node
	node.next.prev = node
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	v, ok := this.hm[key]
	if !ok {
		node := &DlinkNode{val: 1, key: key}
		this.insertBeforeTail(node)
		this.hm[key] = node
	} else {
		v.val++
		for v.val > v.prev.val {
			p := v.prev
			swapAdjacent(p, v)
		}
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	v, ok := this.hm[key]
	if ok {
		if v.val == 1 {
			v.next.prev = v.prev
			v.prev.next = v.next
			v.prev = nil
			v.next = nil
			delete(this.hm, key)
		} else {
			v.val--
			for v.val < v.next.val {
				n := v.next
				swapAdjacent(v, n)
			}
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	return this.head.next.key
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	return this.tail.prev.key
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
