package p324

import (
	"sort"
)

/**
Given an unsorted array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....

Example:
(1) Given nums = [1, 5, 1, 1, 6, 4], one possible answer is [1, 4, 1, 5, 1, 6].
(2) Given nums = [1, 3, 2, 2, 3, 1], one possible answer is [2, 3, 1, 3, 1, 2].

Note:
You may assume all input has valid answer.

Follow Up:
Can you do it in O(n) time and/or in-place with O(1) extra space?

*/

// todo: a faster solution

func wiggleSort(nums []int) {
	sort.Ints(nums)

	half := (len(nums) + 1) / 2
	tmp := make([]int, len(nums)-half)
	copy(tmp, nums[half:])

	c := len(nums) - 1
	for i := 0; i < half; i++ {
		nums[c] = nums[i]
		c--
	}
	c++
	for i := 0; i < len(nums); i += 2 {
		nums[i] = nums[c]
		c++
	}

	c = len(tmp) - 1
	for i := 1; i < len(nums); i += 2 {
		nums[i] = tmp[c]
		c--
	}
}
