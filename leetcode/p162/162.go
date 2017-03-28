package p162

import "fmt"

/**
A peak element is an element that is greater than its neighbors.

Given an input array where num[i] ≠ num[i+1], find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that num[-1] = num[n] = -∞.

For example, in array [1, 2, 3, 1], 3 is a peak element and your function should return the index number 2.
 */

func findPeadElementWithOffset(offset int, nums []int) int {
	fmt.Println(offset, nums)
	length := len(nums)
	if length == 0 {
		return -1
	} else if length == 1 {
		return 0 + offset
	} else if length == 2 {
		if nums[0] > nums[1] {
			return 0 + offset
		} else {
			return 1 + offset
		}
	}
	mean := length / 2
	if nums[mean] > nums[mean-1] && nums[mean] > nums[mean+1] {
		return mean + offset
	} else if nums[mean] > nums[mean+1] {
		return findPeadElementWithOffset(offset, nums[:mean])
	} else {
		return findPeadElementWithOffset(offset+mean, nums[mean:])
	}

}

func findPeakElement(nums []int) int {
	return findPeadElementWithOffset(0, nums)
}
