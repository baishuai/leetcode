package p027

/**
Given an array and a value, remove all instances of that value in place and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example:
Given input array nums = [3,2,2,3], val = 3

Your function should return length = 2, with the first two elements of nums being 2.
 */

func removeElement(nums []int, val int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		if nums[hi] == val {
			hi --
			continue
		}
		if nums[lo] != val {
			lo++
			continue
		}
		nums[lo], nums[hi] = nums[hi], nums[lo]
	}

	if len(nums) <= 0 || nums[0] == val {
		lo--
	}
	return lo + 1
}
