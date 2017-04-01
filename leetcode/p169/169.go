package p169

/**
Given an array of size n, find the majority element.
 The majority element is the element that appears more than ⌊ n/2 ⌋ times.
You may assume that the array is non-empty and the majority element always exist in the array.
 */

func majorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
			count++
		} else {
			if v == candidate {
				count++
			} else {
				count--
			}
		}
	}
	return candidate
}
