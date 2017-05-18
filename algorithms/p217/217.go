package p217

import "sort"

/**
Given an array of integers, find if the array contains any duplicates. Your function should return true if any value appears at least twice in the array,
 and it should return false if every element is distinct.
 */

func containsDuplicate(nums []int) bool {

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}
