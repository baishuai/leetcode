package p081

import "sort"

/**
Follow up for "Search in Rotated Sorted Array":
What if duplicates are allowed?

Would this affect the run-time complexity? How and why?
Suppose an array sorted in ascending order is rotated at some
 pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

Write a function to determine if a given target is in the array.

The array may contain duplicates.
*/

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	lo, hi, mid := 0, len(nums)-1, 0
	for lo < hi-1 {
		mid = (lo + hi) / 2
		if nums[lo] < nums[mid] {
			lo = mid
		} else if nums[lo] > nums[mid] {
			hi = mid
		} else if nums[lo] == nums[lo+1] {
			lo++
		} else if nums[hi] == nums[hi-1] {
			hi--
		}
	}
	lo++
	var ans int
	if target < nums[0] {
		ans = sort.SearchInts(nums[lo:], target) + lo
	} else {
		ans = sort.SearchInts(nums[:lo], target)
	}
	return ans < len(nums) && nums[ans] == target
}
