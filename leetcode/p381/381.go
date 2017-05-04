package p381

import (
	"math/rand"
)

/**
Design a data structure that supports all following operations in average O(1) time.

Note: Duplicate elements are allowed.
insert(val): Inserts an item val to the collection.
remove(val): Removes an item val from the collection if present.
getRandom: Returns a random element from current collection of elements. The probability of each element being returned is linearly related to the number of same value the collection contains.
 */

type valLoc struct {
	val int
	loc int
}

type locsLen struct {
	locs []int
	size int
}

type RandomizedCollection struct {
	vec  []*valLoc
	locs map[int]*locsLen
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		vec:  make([]*valLoc, 0),
		locs: make(map[int]*locsLen),
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	v, ok := this.locs[val]
	if !ok {
		this.locs[val] = &locsLen{locs: make([]int, 0), size: 0}
		v = this.locs[val]
	}
	tail := len(this.vec)
	v.locs = append(v.locs, tail)
	this.vec = append(this.vec, &valLoc{val: val, loc: v.size})
	v.size++
	return !ok
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	v, ok := this.locs[val]
	if ok {
		v.size--
		vpos := v.locs[v.size]
		v.locs = v.locs[:v.size]
		if v.size == 0 {
			delete(this.locs, val)
		}
		if vpos != len(this.vec)-1 {
			tail := this.vec[len(this.vec)-1]
			this.vec[vpos] = tail
			this.locs[tail.val].locs[tail.loc] = vpos
		}
		this.vec = this.vec[:len(this.vec)-1]
	}
	return ok
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.vec[rand.Intn(len(this.vec))].val
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
