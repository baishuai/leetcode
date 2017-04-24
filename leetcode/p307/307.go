package p307

//https://www.hrwhisper.me/binary-indexed-tree-fenwick-tree/
//segment tree

/**
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

The update(i, val) function modifies nums by updating the element at index i to val.
Example:
Given nums = [1, 3, 5]

sumRange(0, 2) -> 9
update(1, 2)
sumRange(0, 2) -> 8
Note:
The array is only modifiable by the update function.
You may assume the number of calls to update and sumRange function is distributed evenly.
*/

type NumArray struct {
	n    int
	segt []int
	nums []int
}

func lowbit(x int) int {
	return x & -x
}

func Constructor(nums []int) NumArray {
	na := NumArray{n: len(nums), segt: make([]int, len(nums)+1), nums: make([]int, len(nums))}
	for i, v := range nums {
		na.add(i+1, v)
	}
	copy(na.nums, nums)
	return na
}

// x start from 1
func (this *NumArray) sum(x int) int {
	res := 0
	for x > 0 {
		res += this.segt[x]
		x -= lowbit(x)
	}
	return res
}

// x start from 1
func (this *NumArray) add(x, val int) {
	for x <= this.n {
		this.segt[x] += val
		x += lowbit(x)
	}
}

func (this *NumArray) Update(i int, val int) {
	diff := val - this.nums[i]
	this.nums[i] = val
	i++
	this.add(i, diff)
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sum(j+1) - this.sum(i)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
