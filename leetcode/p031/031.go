package p031

import (
	"sort"
)

/**
Implement next permutation, which rearranges numbers into the
lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it
as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place, do not allocate extra memory.

Here are some examples. Inputs are in the left-hand column and
its corresponding outputs are in the right-hand column.
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
 */

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	reverse := func(l int) {
		h := len(nums)-1
		for l < h {
			nums[l], nums[h] = nums[h], nums[l]
			l++
			h--
		}
	}

	var j int
	for j = len(nums) - 2; j >= 0; j-- {
		if nums[j] < nums[j+1] {
			//找到刚好大一点的 then swap
			reverse(j + 1)
			find := sort.SearchInts(nums[j+1:], nums[j]+1) + j + 1
			nums[j], nums[find] = nums[find], nums[j]
			break
		}
	}
	if j == -1 {
		reverse(0)
	}

}
