package p380

import "math/rand"

/**
Design a data structure that supports all following operations in average O(1) time.

insert(val): Inserts an item val to the set if not already present.
remove(val): Removes an item val from the set if present.
getRandom: Returns a random element from current set of elements.
Each element must have the same probability of being returned.
 */

type RandomizedSet struct {
	arr []int
	set map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		arr: make([]int, 0),
		set: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; !ok {
		this.set[val] = len(this.arr)
		this.arr = append(this.arr, val)
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.set[val]; ok {
		tail := this.arr[len(this.arr)-1]
		if tail != val {
			this.arr[index] = tail
			this.set[tail] = index
		}
		delete(this.set, val)
		this.arr = this.arr[:len(this.arr)-1]
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]

}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
