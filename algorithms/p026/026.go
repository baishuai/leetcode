package p026

/**
Given a sorted array, remove the duplicates in place such that each element
 appear only once and return the new length.

Do not allocate extra space for another array, you must do this in place
 with constant memory.

For example,
Given input array nums = [1,1,2],

Your function should return length = 2, with the first two elements
 of nums being 1 and 2 respectively.
It doesn't matter what you leave beyond the new length.
*/

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	first := nums[0]
	length := 1
	index := 1
	for _, v := range nums[1:] {
		if v != first {
			first = v
			nums[index] = v
			index++
			length++
		}
	}
	return length
}
