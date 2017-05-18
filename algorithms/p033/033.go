package p033

import (
	"sort"
)

/**
Suppose an array sorted in ascending order is rotated at
 some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

You are given a target value to search. If found in the array
 return its index, otherwise return -1.

You may assume no duplicate exists in the array.
 */

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi, mid := 0, len(nums)-1, 0
	for lo < hi-1 {
		mid = (lo + hi) / 2
		if nums[lo] < nums[mid] {
			lo = mid
		} else if nums[lo] > nums[mid] {
			hi = mid
		}
	}
	lo++
	var ans int
	if target < nums[0] {
		ans = sort.SearchInts(nums[lo:], target) + lo
	} else {
		ans = sort.SearchInts(nums[:lo], target)
	}
	if ans == len(nums) || nums[ans] != target {
		ans = -1
	}
	return ans
}
