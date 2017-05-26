package p581

import (
	"sort"
)

/**
Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order, too.

You need to find the shortest such subarray and output its length.

Example 1:
Input: [2, 6, 4, 8, 10, 9, 15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
Note:
Then length of the input array is in range [1, 10,000].
The input array may contain duplicates, so ascending order here means <=.
**/

// Max return the max value of a and b
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min return the min value of a and b
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findUnsortedSubarray(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return 0
	}
	start, end := -1, -2
	max, min := nums[0], nums[len(nums)-1]
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		max = Max(max, nums[i])
		if nums[i] < max {
			end = i
		}
		min = Min(min, nums[j])
		if nums[j] > min {
			start = j
		}
	}
	return end - start + 1
}

func findUnsortedSubarray2(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return 0
	}
	leftv, lefti, order := nums[0], 0, true
	for i, v := range nums {
		if v < leftv {
			order = false
			lefti = sort.SearchInts(nums[:lefti+1], v+1)
			leftv = v
		} else if order {
			lefti = i
			leftv = v
		}
	}

	rightv, righti, order := nums[len(nums)-1], len(nums)-1, true
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > rightv {
			order = false
			s := nums[righti:len(nums)]
			righti = sort.Search(len(s), func(j int) bool { return s[j] >= nums[i] }) + righti
			rightv = nums[i]
		} else if order {
			righti = i
			rightv = nums[i]
		}
	}
	if righti < lefti {
		return 0
	}
	return righti - lefti
}
