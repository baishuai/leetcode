package p153

/**
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

Find the minimum element.

You may assume no duplicate exists in the array.
*/

func findMin(nums []int) int {

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
	min := nums[0]
	if lo < len(nums) && nums[lo] < nums[0] {
		min = nums[lo]
	}
	return min
}
